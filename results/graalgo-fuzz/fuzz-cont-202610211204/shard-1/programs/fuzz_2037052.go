package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint16 = 255
    var v2 float64 = 1e10
    v3 := []uint32{816}
    var v4 float64 = v2
    for fi5 := 0; fi5 < 19; fi5++ {
        v1 -= v1
        switch -184 {
        case 0:
            var v6 int64 = -1
            _ = v6
        default:
            if (("a" != "a") || (!true)) {
                v7 := GenS1{F0: -1}
                v2 = v4
                _ = v7
            }
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 float32 = (3.14159 + (-2.25 - (3.14159 * 1e10)))
        _ = fi8
        _ = v9
    }
    var v10 uint = 0
    var v11 int16 = 32767
    var v12 int16 = 0
    if false {
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS1{F0: int8(v1)}
            _ = fi13
            _ = v14
        }
        switch -9223372036854775808 {
        case 0:
            var v15 string = (("a" + "") + "a")
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch "hello" {
            case "foo bar":
                v11 = -243
            case "":
                v1 += 662
            default:
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            }
            _ = fi16
        }
    }
    v4 += 3.14159
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 uint = 699
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 string = "foo bar"
    fmt.Println(v11, v10)
    v21 := GenS2{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
