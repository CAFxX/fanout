package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint16 = 1
    var v2 int16 = 42
    var v3 int32 = 0
    if ((!true) && true) {
        if true {
            var v4 int32 = (v3 - (v3 ^ v3))
            v5 := GenS1{F0: 1}
            _ = v4
            _ = v5
        }
        var v6 int16 = (-(v2))
        _ = v6
    }
    var v7 int8 = -1
    var v8 float32 = (3.14159 + 0.1)
    var v9 float64 = (((((float64(v8) / 1.5) - 3.14159) - (0.0 + 3.14159)) / float64(3.14159)) - 1.5)
    if true {
        v7 = (v7 / ((-1) | 1))
    } else {
        fmt.Println(v3)
        v10 := GenS1{F0: -187}
        fmt.Println(v9, v7)
        _ = v10
    }
    v11 := []int{1, -924, -214}
    v12 := GenS1{F0: (v3 | v3)}
    if false {
        var v13 int64 = -9223372036854775808
        _ = v13
    } else {
        v11 = append(v11, 32768)
        v14 := GenS2{F0: (^(v7)), F1: 42}
        switch ("\n\t" + "foo bar") {
        case "hello":
            if false {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v15 := GenS1{F0: v3}
                for fi16 := 0; fi16 < 2; fi16++ {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        var v18 uint = 65535
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                _ = v15
            } else {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 int32 = v3
                    _ = fi19
                    _ = v20
                }
                var v21 uint = 1
                fmt.Println(v1)
                _ = v21
            }
        }
        _ = v14
    }
    if true {
        v11 = append(v11, 127)
    }
    var v22 bool = (!false)
    var v23 uint = 18446744073709551615
    v24 := GenS2{F0: int8(v2), F1: 77}
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS3{F0: v3}
        _ = fi25
        _ = v26
    }
    if len(v11) > 0 {
        v11[0] = -454
    }
    v7 += (v7 << ((v23) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
