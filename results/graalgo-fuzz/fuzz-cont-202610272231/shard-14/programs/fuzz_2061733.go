package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = false
    var v2 bool = v1
    var v3 float32 = 3.14159
    v4 := make(map[uint32]bool)
    v5 := []uint32{1, 886, 782}
    v6 := []int{9223372036854775807}
    var v7 uint16 = 1
    var v8 int32 = -1
    for fi9 := 0; fi9 < 13; fi9++ {
        switch (fi9 >> ((255) & 63)) {
        case 9223372036854775807:
            if (0 == -1) {
                v10 := GenS1{F0: "foo bar"}
                switch ("" + "\n\t") {
                case "\n\t":
                    v11 := GenS1{F0: ("" + "")}
                    _ = v11
                case "a":
                    if true {
                        v2 = (32767 == v7)
                    }
                default:
                    v12 := GenS1{F0: ("hello" + ((("foo bar" + ("\n\t" + "")) + "a") + (("a" + "a") + "\n\t")))}
                    _ = v12
                }
                v5 = append(v5, 2147483648)
                _ = v10
            }
        case 0:
            if len(v6) > 0 {
                v6[0] = -915
            }
        case 32767:
            var v13 int16 = 406
            _ = v13
        }
        v14 := v4[42]
        var v15 uint64 = 1
        _ = fi9
        _ = v14
        _ = v15
    }
    var v16 int32 = v8
    v16 = 1
    v1 = ("hello" <= "foo bar")
    v2 = true
    v17 := GenS2{F0: 0}
    if false {
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 uint32 = 65535
            v2 = (false && v2)
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            _ = fi18
            _ = v19
        }
    }
    if false {
        var v20 bool = false
        _ = v20
    } else {
        var v21 string = "a"
        _ = v21
    }
    v6 = append(v6, -9223372036854775808)
    var v22 uint8 = 42
    if v2 {
        _ = len(v5)
    } else {
        if true {
            v23 := GenS2{F0: -1}
            _ = v23
        } else {
            for fi24 := 0; fi24 < 6; fi24++ {
                v3 = -2.25
                _ = fi24
            }
        }
    }
    fmt.Println(v8, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
