package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint64 {
    var v5 uint8 = 0
    if ((!(32767 <= 32767)) && false) {
        var v6 uint16 = 901
        _ = v6
    }
    fmt.Println(v5)
    v5 = (^(v5))
    _ = v5
    return 42
}

func main() {
    v1 := []string{"foo bar", ""}
    var v2 uint8 = 1
    var v3 uint16 = 688
    var v4 uint = 255
    switch -468 {
    case -2147483648:
        v3 -= (v3 * (^(v3)))
    case 32768:
        v3 = v3
    case -1:
        if true {
            var v7 string = ("foo bar" + "a")
            _ = v7
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS1{F0: 0}
    switch 128 {
    case -177:
        v2 = v2
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if (1000 > 42) {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        } else {
            for fi9 := 0; fi9 < 9; fi9++ {
                v4 -= 2
                _ = fi9
            }
        }
    }
    var v10 uint32 = 2
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 string = (("foo bar" + "a") + (("" + "") + ("a" + ((("a" + "foo bar") + "hello") + "\n\t"))))
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint32 = (v10 << ((18446744073709551615) & 63))
            var v15 int64 = -464
            v4 = uint(v2)
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 string = ("\n\t" + v12)
        _ = fi11
        _ = v12
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 bool = (((!false) && false) && false)
        v3 = v3
        _ = fi17
        _ = v18
    }
    v10 -= v10
    if true {
        v3 += (^(v3))
    } else {
        v10 = v10
        fmt.Println(v2, v3, v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
