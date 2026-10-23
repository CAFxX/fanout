package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }

func main() {
    var v1 int32 = 42
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 int = -1
    if true {
        var v5 uint8 = 128
        delete(v2, ((("foo bar" + "\n\t") + ("" + ((("hello" + "\n\t") + (("\n\t" + "\n\t") + "foo bar")) + "\n\t"))) + "hello"))
        fmt.Println(v4, v1)
        _ = v5
    }
    v1 = v1
    if false {
        var v6 int32 = v1
        switch (v4 >> ((441) & 63)) {
        case 2147483647:
            var v7 float64 = 1.5
            _ = v7
        }
        var v8 uint8 = 42
        _ = v6
        _ = v8
    }
    if (0 != -706) {
        v4 = 32767
    } else {
        v9 := GenS1{F0: 0, F1: -1}
        _ = v9
    }
    v10 := GenS1{F0: 867, F1: 42}
    var v11 int = 0
    if true {
        var v12 int16 = -32768
        _ = v12
    } else {
        if (70 == int8(v1)) {
            v13 := GenS1{F0: 42, F1: 32767}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint8 = 2
                var v16 uint8 = v15
                v11 = fi14
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        v17 := GenS1{F0: -443, F1: -32768}
        v18 := GenS1{F0: -32768, F1: 0}
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 uint64 = 2
        _ = fi19
        _ = v20
    }
    delete(v2, "a")
    var v21 uint16 = 255
    v22 := GenS1{F0: int16(v4), F1: 577}
    v21 += 460
    var v23 float64 = 3.14159
    v24 := []int{(v11 % ((v11) | 1)), (v4 % ((v11) | 1))}
    var v25 int8 = 127
    v1 = int32(v21)
    v26 := GenS1{F0: 42, F1: -1}
    var v27 int8 = (v25 % ((-1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
