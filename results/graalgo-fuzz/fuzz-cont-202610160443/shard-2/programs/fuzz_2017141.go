package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 int32, p1 uint8) uint32 {
    v6 := []int{}
    _ = len(v6)
    for fi7 := 0; fi7 < 4; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    p0 += p0
    var v8 bool = ((3.14159 <= -2.25) || (0 >= 1))
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v4 int8 = -1
    var v5 int16 = 32767
    var v9 int32 = 2147483647
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := v2[0]
        switch "a" {
        case "hello":
            if v11 {
                v12 := []int{fi10, (fi10 & fi10)}
                switch (^(fi10)) {
                case -2147483648:
                    if (442 <= 0) {
                        var v13 int8 = (v4 | 100)
                        _ = v13
                    } else {
                        v11 = (18446744073709551615 < 4294967295)
                        fmt.Println(v9, v11)
                    }
                default:
                    v14 := GenS1{F0: 2, F1: 115}
                    _ = v14
                }
                var v15 int16 = v5
                _ = v12
                _ = v15
            }
        }
        switch ("foo bar" + "hello") {
        case "\n\t":
            if (0 == 790) {
                fmt.Println(v9)
                fmt.Println(v11)
                v9 += (-(v9))
            }
        case "":
            switch "a" {
            case "\n\t":
                v9 = -106
            case "":
                v11 = (128 > 127)
            }
        }
        _ = fi10
        _ = v11
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 bool = (false || true)
        var v18 uint64 = 18446744073709551615
        v19 := GenS1{F0: 784, F1: 924}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            _ = len(v1)
            v22 := []int{}
            delete(v2, 309)
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS2{F0: 253, F1: (^(fi23)), F2: 329}
            _ = fi23
            _ = v24
        }
        v25 := GenS2{F0: 42, F1: fi20, F2: 111}
        _ = fi20
        _ = v25
    }
    var v26 uint32 = 1
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 uint8 = 127
        var v29 float32 = -2.25
        var v30 uint = 373
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    v31 := []int{}
    for fi32 := 0; fi32 < 4; fi32++ {
        v4 += 127
        _ = fi32
    }
    var v33 uint32 = v26
    var v34 string = ""
    v33 += (v33 ^ v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
