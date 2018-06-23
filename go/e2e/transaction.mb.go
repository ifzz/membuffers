package e2e

import (
	"github.com/orbs-network/membuffers/go"
)

/*
message Transaction {
	TransactionData data = 1;
	bytes signature = 2;
}
*/

// reader

type Transaction struct {
	_Message membuffers.Message
}

var m_Transaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes}
var m_Transaction_Unions = [][]membuffers.FieldType{{}}

func ReadTransaction(buf []byte) *Transaction {
	x := &Transaction{}
	x._Message.Init(buf, membuffers.Offset(len(buf)), m_Transaction_Scheme, m_Transaction_Unions)
	return x
}

func (x *Transaction) _RawBuffer() []byte {
	return x._Message.RawBuffer()
}

func (x *Transaction) Data() *TransactionData {
	b, s := x._Message.GetMessage(0)
	return ReadTransactionData(b[:s])
}

func (x *Transaction) _RawBuffer_Data() []byte {
	return x._Message.RawBufferForField(0, 0)
}

func (x *Transaction) Signature() []byte {
	return x._Message.GetBytes(1)
}

func (x *Transaction) _RawBuffer_Signature() []byte {
	return x._Message.RawBufferForField(1, 0)
}

// writer

type TransactionWriter struct {
	_Writer membuffers.Writer
	Data *TransactionDataWriter
	Signature []byte
}

func (w *TransactionWriter) Write(buf []byte) {
	w._Writer.Reset()
	w._Writer.WriteMessage(buf, w.Data)
	w._Writer.WriteBytes(buf, w.Signature)
}

func (w *TransactionWriter) GetSize() membuffers.Offset {
	return w._Writer.GetSize()
}

/*
message TransactionData {
	uint32 protocol_version = 1;
	uint64 virtual_chain = 2;
	repeated TransactionSender sender = 3;
	int64 time_stamp = 4;
}
*/

// reader

type TransactionData struct {
	_Message membuffers.Message
}

var m_TransactionData_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessageArray,membuffers.TypeUint64}
var m_TransactionData_Unions = [][]membuffers.FieldType{{}}

func ReadTransactionData(buf []byte) *TransactionData {
	x := &TransactionData{}
	x._Message.Init(buf, membuffers.Offset(len(buf)), m_TransactionData_Scheme, m_TransactionData_Unions)
	return x
}

func (x *TransactionData) _RawBuffer() []byte {
	return x._Message.RawBuffer()
}

func (x *TransactionData) ProtocolVersion() uint32 {
	return x._Message.GetUint32(0)
}

func (x *TransactionData) _RawBuffer_ProtocolVersion() []byte {
	return x._Message.RawBufferForField(0, 0)
}

func (x *TransactionData) VirtualChain() uint64 {
	return x._Message.GetUint64(1)
}

func (x *TransactionData) _RawBuffer_VirtualChain() []byte {
	return x._Message.RawBufferForField(1, 0)
}

func (x *TransactionData) SenderIterator() *TransactionData_SenderIterator {
	return &TransactionData_SenderIterator{_Iterator: x._Message.GetMessageArrayIterator(2)}
}

type TransactionData_SenderIterator struct {
	_Iterator *membuffers.Iterator
}

func (i *TransactionData_SenderIterator) HasNext() bool {
	return i._Iterator.HasNext()
}

func (i *TransactionData_SenderIterator) NextSender() *TransactionSender {
	b, s := i._Iterator.NextMessage()
	return ReadTransactionSender(b[:s])
}

func (x *TransactionData) _RawBuffer_SenderArray() []byte {
	return x._Message.RawBufferForField(2, 0)
}

func (x *TransactionData) TimeStamp() uint64 {
	return x._Message.GetUint64(3)
}

func (x *TransactionData) _RawBuffer_TimeStamp() []byte {
	return x._Message.RawBufferForField(3, 0)
}

// writer

type TransactionDataWriter struct {
	_Writer membuffers.Writer
	ProtocolVersion uint32
	VirtualChain uint64
	Sender []*TransactionSenderWriter
	TimeStamp uint64
}

func (w *TransactionDataWriter) sender() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Sender))
	for i, v := range w.Sender {
		res[i] = v
	}
	return res
}

func (w *TransactionDataWriter) Write(buf []byte) {
	w._Writer.Reset()
	w._Writer.WriteUint32(buf, w.ProtocolVersion)
	w._Writer.WriteUint64(buf, w.VirtualChain)
	w._Writer.WriteMessageArray(buf, w.sender())
	w._Writer.WriteUint64(buf, w.TimeStamp)
}

func (w *TransactionDataWriter) GetSize() membuffers.Offset {
	return w._Writer.GetSize()
}

/*
message TransactionSender {
	string name = 1;
	repeated string friend = 2;
}
*/

// reader

type TransactionSender struct {
	_Message membuffers.Message
}

var m_TransactionSender_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeStringArray}
var m_TransactionSender_Unions = [][]membuffers.FieldType{{}}

func ReadTransactionSender(buf []byte) *TransactionSender {
	x := &TransactionSender{}
	x._Message.Init(buf, membuffers.Offset(len(buf)), m_TransactionSender_Scheme, m_TransactionSender_Unions)
	return x
}

func (x *TransactionSender) _RawBuffer() []byte {
	return x._Message.RawBuffer()
}

func (x *TransactionSender) Name() string {
	return x._Message.GetString(0)
}

func (x *TransactionSender) _RawBuffer_Name() []byte {
	return x._Message.RawBufferForField(0, 0)
}

func (x *TransactionSender) FriendIterator() *TransactionSender_FriendIterator {
	return &TransactionSender_FriendIterator{_Iterator: x._Message.GetStringArrayIterator(1)}
}

type TransactionSender_FriendIterator struct {
	_Iterator *membuffers.Iterator
}

func (i *TransactionSender_FriendIterator) HasNext() bool {
	return i._Iterator.HasNext()
}

func (i *TransactionSender_FriendIterator) NextFriend() string {
	return i._Iterator.NextString()
}

func (x *TransactionSender) _RawBuffer_FriendArray() []byte {
	return x._Message.RawBufferForField(1, 0)
}

// writer

type TransactionSenderWriter struct {
	_Writer membuffers.Writer
	Name string
	Friend []string
}

func (w *TransactionSenderWriter) Write(buf []byte) {
	w._Writer.Reset()
	w._Writer.WriteString(buf, w.Name)
	w._Writer.WriteStringArray(buf, w.Friend)
}

func (w *TransactionSenderWriter) GetSize() membuffers.Offset {
	return w._Writer.GetSize()
}