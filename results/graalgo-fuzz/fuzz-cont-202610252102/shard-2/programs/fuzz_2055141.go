package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
func genFn0() string {
    if false {
        v6 := GenS1{F0: 9223372036854775807, F1: 140, F2: ("hello" + "foo bar")}
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            var v8 uint8 = 113
            var v9 uint64 = 18446744073709551615
            v10 := []int{861}
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 float64 = 1.5
                _ = fi11
                _ = v12
            }
            switch (("" + ("a" + ("a" + ("\n\t" + "hello")))) + "hello") {
            case "\n\t":
                v13 := []int{(-(fi7)), fi7, fi7, fi7}
                _ = v13
            case "":
                v14 := GenS1{F0: 9223372036854775807, F1: 42, F2: "foo bar"}
                _ = v14
            default:
                var v15 string = "a"
                _ = v15
            }
        }
        _ = fi7
    }
    var v16 uint32 = 65535
    for fi17 := 0; fi17 < 11; fi17++ {
        v16 = 42
        var v18 uint8 = 1
        if (18446744073709551615 != 1) {
            v18 += (v18 << ((42) & 63))
        } else {
            v18 += 42
            for fi19 := 0; fi19 < 11; fi19++ {
                v16 -= (v16 * (v16 | 381))
                var v20 float32 = float32(((3.14159 + float64(1.5)) - ((1.5 - 1.5) + 0.0)))
                v21 := []int{int(v16)}
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        _ = fi17
        _ = v18
    }
    var v22 bool = true
    _ = v16
    _ = v22
    return ("hello" + "a")
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    var v3 uint64 = 227
    v4 := make(map[string]uint32)
    var v5 int16 = -459
    var v23 int32 = 0
    var v24 int64 = -1
    var v25 uint16 = 32767
    fmt.Println(v5)
    fmt.Println(v23)
    var v26 int = 9223372036854775807
    _ = len(v4)
    v24 = v24
    if (false && true) {
        _ = len(v4)
        v27 := GenS1{F0: v24, F1: 0, F2: ""}
        _ = v27
    }
    v5 = v5
    var v28 int16 = (v5 << ((v1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
