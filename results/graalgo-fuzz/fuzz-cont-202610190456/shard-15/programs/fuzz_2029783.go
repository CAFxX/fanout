package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint64) uint {
    var v6 float64 = 1.5
    for fi7 := 0; fi7 < 14; fi7++ {
        switch ("a" + "foo bar") {
        case "":
            if (true || (-747 != 511)) {
                var v8 int16 = -32768
                _ = v8
            } else {
                var v9 string = ((("" + "hello") + "a") + ("\n\t" + ("" + "")))
                _ = v9
            }
        default:
            fmt.Println(v6, p0)
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := []string{"a", "\n\t", "\n\t"}
    v2 := make(map[string]bool)
    var v3 string = ""
    var v4 uint64 = 18446744073709551615
    var v5 uint8 = 42
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v10 uint32 = 42
    fmt.Println(v10)
    fmt.Println(v3, v4)
    _ = len(v2)
    var v11 bool = true
    fmt.Println(v4, v11)
    if v11 {
        v4 = 384
        v11 = v11
    }
    fmt.Println(v11, v4)
    switch -721 {
    case 0:
        v12 := GenS1{F0: v4}
        _ = v12
    case -205:
        if len(v1) > 0 {
            v1[0] = "a"
        }
    default:
        v10 += v10
    }
    switch -263 {
    case 801:
        switch 32768 {
        case 32768:
            var v13 uint16 = 42
            _ = v13
        }
    default:
        fmt.Println(v10)
    }
    _ = len(v1)
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int32 = -1
        _ = fi14
        _ = v15
    }
    var v16 uint32 = v10
    if true {
        switch 870 {
        case -1:
            for fi17 := 0; fi17 < 19; fi17++ {
                if true {
                    var v18 uint8 = (v5 << ((0) & 63))
                    _ = v18
                } else {
                    v5 = 5
                }
                _ = fi17
            }
        case 445:
            v11 = false
        }
        v5 = 42
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 int8 = 42
            v21 := GenS1{F0: uint64(v10)}
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        fmt.Println(v11, v4, v11)
        var v22 string = v3
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
