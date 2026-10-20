package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 bool }
func genFn0() int8 {
    v5 := GenS1{F0: 2, F1: 42, F2: 1}
    for fi6 := 0; fi6 < 11; fi6++ {
        if ((!false) && false) {
            fmt.Println(fi6)
            var v7 uint16 = 2
            _ = v7
        }
        v8 := []int{(fi6 + fi6)}
        _ = fi6
        _ = v8
    }
    v9 := []int{}
    switch ("\n\t" + (("" + ("\n\t" + "foo bar")) + "\n\t")) {
    case "hello":
        switch 2147483647 {
        case 1:
            v10 := GenS1{F0: 661, F1: 59, F2: 255}
            _ = v10
        case 32767:
            switch ("foo bar" + "a") {
            case "":
                for fi11 := 0; fi11 < 6; fi11++ {
                    v12 := GenS1{F0: 255, F1: 11, F2: 32767}
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    v13 := GenS1{F0: 162, F1: 128, F2: 255}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case "a":
                _ = len(v9)
            case "\n\t":
                v14 := GenS1{F0: 1, F1: 255, F2: 65535}
                _ = v14
            default:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
        }
    default:
        var v15 string = "\n\t"
        _ = v15
    }
    if len(v9) > 0 {
        v9[0] = 32768
    }
    _ = v5
    _ = v9
    return 42
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 int8 = 1
    fmt.Println(v4, v1)
    fmt.Println(v1, v4)
    var v16 bool = (true && (956 != 816))
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: v1, F1: 2, F2: uint16(v1)}
        v4 -= genFn0()
        v1 += (v1 & (v1 * v1))
        _ = fi17
        _ = v18
    }
    delete(v2, 42)
    v1 = v1
    fmt.Println(v16)
    fmt.Println(v4)
    var v19 uint32 = 65535
    if v16 {
        v20 := GenS1{F0: 18446744073709551615, F1: 168, F2: 2}
        var v21 uint8 = 2
        _ = v20
        _ = v21
    }
    var v22 float64 = -2.25
    v23 := v2[(v19 - 248)]
    fmt.Println(v1, v23, v4)
    var v24 int32 = -1
    if v16 {
        var v25 uint32 = 0
        v26 := []int{(-(v23)), v23, (v23 / (((v23 + 32767)) | 1))}
        v27 := GenS2{F0: (v25 & v25), F1: 4294967295, F2: v16}
        _ = v25
        _ = v26
        _ = v27
    }
    v1 = v1
    var v28 uint32 = (v19 << ((v1) & 63))
    var v29 int8 = (v4 / (((^(v4))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
