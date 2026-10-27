package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint) uint {
    v6 := GenS1{F0: 370}
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(p0, p0)
        p0 = p0
        _ = fi7
    }
    _ = p0
    _ = v6
    return (p0 - p0)
}
func genFn1() uint32 {
    var v8 uint = 2
    var v9 uint = 42
    _ = v8
    _ = v9
    return 407
}

func main() {
    v1 := []string{((("\n\t" + "\n\t") + "") + "a"), "a", ("" + "foo bar"), ""}
    var v2 int8 = 0
    var v3 uint64 = 1
    v4 := make(map[uint32]int)
    var v5 uint = 255
    if (!true) {
        var v10 uint16 = 2
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint32 = 65535
            _ = fi11
            _ = v12
        }
        var v13 int16 = 32767
        _ = v10
        _ = v13
    }
    var v14 bool = ((212 != 32767) || (32767 <= -996))
    v2 = 100
    v4[0] = 1
    var v15 uint = v5
    v15 += 0
    for fi16 := 0; fi16 < 5; fi16++ {
        if v14 {
            v17 := GenS1{F0: int16(fi16)}
            _ = v17
        } else {
            switch "\n\t" {
            case "foo bar":
                _ = len(v1)
            }
            v18 := GenS1{F0: 42}
            _ = v18
        }
        if v14 {
            var v19 uint8 = 4
            v20 := GenS1{F0: 0}
            var v21 float32 = 0.1
            _ = v19
            _ = v20
            _ = v21
        }
        if v14 {
            var v22 int = fi16
            var v23 uint32 = 1
            _ = v22
            _ = v23
        }
        _ = fi16
    }
    var v24 uint8 = 1
    v5 = (v5 % (((v5 ^ 4294967295)) | 1))
    for fi25 := 0; fi25 < 9; fi25++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v14)
        _ = fi25
    }
    fmt.Println(v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
