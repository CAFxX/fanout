package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }

func main() {
    var v1 uint32 = 0
    var v2 int8 = 127
    v3 := []string{"foo bar", ""}
    var v4 float32 = (0.1 - 0.0)
    var v5 int32 = -1
    for fi6 := 0; fi6 < 9; fi6++ {
        v4 = v4
        _ = fi6
    }
    fmt.Println(v5)
    v4 -= float32(0.0)
    var v7 int8 = (^(v2))
    v8 := GenS1{F0: 568, F1: 0, F2: ((v2 % ((v2) | 1)) == (v7 >> ((4294967295) & 63)))}
    if (32767 != 330) {
        var v9 uint32 = 65535
        _ = v9
    } else {
        v3 = append(v3, "a")
        v10 := GenS1{F0: 32767, F1: 623, F2: false}
        for fi11 := 0; fi11 < 19; fi11++ {
            switch fi11 {
            case 0:
                v3 = append(v3, "hello")
            case 127:
                var v12 float32 = (1.5 - v4)
                _ = v12
            default:
                v13 := GenS1{F0: 2, F1: 42, F2: (!(942 == 65535))}
                _ = v13
            }
            _ = fi11
        }
        _ = v10
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v14 uint8 = 255
    if true {
        switch (("a" + "foo bar") + "a") {
        case "":
            if (!("hello" > ((("\n\t" + "foo bar") + ("" + ("hello" + ("a" + "\n\t")))) + "hello"))) {
                var v15 int8 = v7
                if len(v3) > 0 {
                    v3[0] = ""
                }
                _ = v15
            } else {
                v16 := GenS1{F0: 255, F1: -1, F2: ("hello" >= "a")}
                for fi17 := 0; fi17 < 17; fi17++ {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 float64 = 3.14159
                        var v20 uint16 = 930
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = fi17
                }
                _ = v16
            }
        }
        var v21 int8 = (v2 % ((v2) | 1))
        _ = v21
    }
    var v22 uint64 = 674
    var v23 bool = false
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(fi24)
        v7 += v2
        _ = fi24
    }
    v3 = append(v3, "\n\t")
    v14 = v14
    for fi25 := 0; fi25 < 16; fi25++ {
        v3 = append(v3, "\n\t")
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := GenS1{F0: 42, F1: -976, F2: v23}
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
