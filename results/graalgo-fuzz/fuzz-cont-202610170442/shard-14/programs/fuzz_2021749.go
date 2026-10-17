package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }
func genFn0() int8 {
    v7 := []int{442, -361, -709}
    var v8 uint8 = 255
    v9 := GenS1{F0: true, F1: 65535, F2: 0}
    _ = len(v7)
    var v10 uint64 = 0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 bool = (true && true)
    v2 := []string{"", ("foo bar" + "hello")}
    var v3 string = ""
    v4 := make(map[string]int)
    var v5 bool = (42 < 42)
    var v6 uint = 332
    var v11 uint = v6
    v12 := v4[("foo bar" + "foo bar")]
    v3 = v3
    var v13 int32 = -1
    v11 += v11
    v14 := v4["a"]
    switch v12 {
    case -1:
        v13 = 42
    case 9223372036854775807:
        var v15 string = "foo bar"
        _ = v15
    case 664:
        var v16 int16 = 32767
        _ = v16
    default:
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS1{F0: v5, F1: 42, F2: 229}
            if v1 {
                if v1 {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        fmt.Println(v14)
                        var v20 int8 = -71
                        var v21 int8 = genFn0()
                        if len(v2) > 0 {
                            v2[0] = (v3 + (v3 + ((v3 + v3) + v3)))
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                } else {
                    var v22 int16 = 769
                    _ = v22
                }
            }
            v23 := v4[(("" + "a") + "hello")]
            _ = fi17
            _ = v18
            _ = v23
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v24 int = v12
    switch (((v3 + ("foo bar" + v3)) + ("\n\t" + v3)) + (v3 + (("foo bar" + v3) + (v3 + v3)))) {
    case "":
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 uint32 = 2147483647
            var v27 int16 = 32767
            v14 = -2147483648
            _ = fi25
            _ = v26
            _ = v27
        }
    case "\n\t":
        _ = len(v4)
    }
    var v28 int8 = -55
    v24 = (v12 << ((v6) & 63))
    fmt.Println(v28, v3, v6)
    var v29 int = v12
    v30 := GenS2{F0: v11, F1: (v28 | v28), F2: 798}
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 string = v3
        _ = fi31
        _ = v32
    }
    var v33 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
