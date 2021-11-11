using System;
using System.Collections.Generic;
using System.Text;

namespace DataStructure
{

    public class MyNode<T>
    {
        public MyNode()
        {

        }

        public MyNode(T value)
        {
            Value = value;
        }

        public MyNode<T> Next { get; set; }

        public T Value { get; set; }

        public MyNode<T> Clone()
        {
            var copyNode = new MyNode<T>();
            copyNode.Value = Value;
            copyNode.Next = Next;
            return copyNode;
        }

        public void Clear()
        {
            Value = default(T);
            Next = default(MyNode<T>);
        }
    }
}
