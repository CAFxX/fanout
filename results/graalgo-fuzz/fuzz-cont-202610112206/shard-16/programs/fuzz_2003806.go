package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }

func main() {
    var v1 bool = ((!(!true)) || true)
    v2 := make(map[uint32]int)
    var v3 int = 4294967295
    v1 = (26 == 42)
    v4 := GenS1{F0: 18446744073709551615, F1: ("a" + "a"), F2: 4294967295}
    v5 := GenS1{F0: 255, F1: "hello", F2: 42}
    v6 := GenS1{F0: 255, F1: ("a" + ("foo bar" + ("\n\t" + "foo bar"))), F2: 18446744073709551615}
    v7 := GenS1{F0: 29, F1: ("\n\t" + "hello"), F2: 0}
    var v8 uint16 = 42
    var v9 uint32 = 2
    _ = len(v2)
    v10 := GenS1{F0: 18446744073709551615, F1: ("hello" + ("\n\t" + "foo bar")), F2: 42}
    for fi11 := 0; fi11 < 17; fi11++ {
        if false {
            var v12 float32 = 1.5
            _ = v12
        }
        var v13 uint8 = 0
        v14 := []int{}
        _ = fi11
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 42, F1: ((("\n\t" + "a") + (("foo bar" + "hello") + "\n\t")) + "hello"), F2: 42}
    fmt.Println(v9, v8)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
