package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float64 = 0.1
    var v2 string = ("foo bar" + "a")
    var v3 int16 = -874
    var v4 uint64 = 4294967295
    if (!(255 > 0)) {
        if true {
            v4 += (v4 * (^(v4)))
        }
    }
    fmt.Println(v3, v2)
    switch v2 {
    case "\n\t":
        switch 128 {
        case 2147483647:
            for fi5 := 0; fi5 < 4; fi5++ {
                for fi6 := 0; fi6 < 10; fi6++ {
                    var v7 int = fi6
                    var v8 float32 = float32(v1)
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                _ = fi5
            }
        }
    case "":
        v4 = (v4 / (((v4 - v4)) | 1))
    case "foo bar":
        v9 := []int{9223372036854775807, 962, -2147483648, 42}
        _ = v9
    default:
        var v10 int32 = -1
        _ = v10
    }
    fmt.Println(v4)
    for fi11 := 0; fi11 < 2; fi11++ {
        switch (-(fi11)) {
        case -596:
            v3 += -32768
        }
        if (607 < -949) {
            var v12 uint64 = v4
            v13 := []int{(-(fi11))}
            if false {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                if len(v13) > 0 {
                    v13[0] = 526
                }
            } else {
                v3 = v3
            }
            _ = v12
            _ = v13
        } else {
            var v14 uint8 = 51
            _ = v14
        }
        v15 := GenS1{F0: 0}
        _ = fi11
        _ = v15
    }
    var v16 uint64 = v4
    for fi17 := 0; fi17 < 20; fi17++ {
        v16 += v4
        v1 -= v1
        var v18 uint64 = (v4 - (^(v16)))
        _ = fi17
        _ = v18
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
