package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 bool = true
    v2 := []string{(("\n\t" + "foo bar") + ("a" + ("\n\t" + "a")))}
    v3 := make(map[uint32]int)
    var v4 int16 = 1
    switch 0 {
    case 1:
        fmt.Println(v4, v4)
    case 0:
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint32 = 583
            delete(v3, v6)
            var v7 string = "\n\t"
            _ = fi5
            _ = v6
            _ = v7
        }
    case 9223372036854775807:
        v4 += (v4 | v4)
    default:
        v3[65535] = 1
    }
    var v8 uint32 = 2
    fmt.Println(v8, v4)
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    _ = len(v3)
    fmt.Println(v4)
    fmt.Println(v8)
    if true {
        fmt.Println(v4, v8, v4)
    }
    v9 := GenS1{F0: -2147483648}
    delete(v3, (^(v8)))
    v1 = (((("" + "\n\t") > (("foo bar" + "\n\t") + "\n\t")) && (v4 <= (v4 | (v4 << ((4294967295) & 63))))) && v1)
    var v10 int64 = 0
    if false {
        v10 -= 0
        if false {
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(fi11)
                _ = fi11
            }
            switch "" {
            case "hello":
                fmt.Println(v1, v4)
            case "":
                var v12 uint64 = 80
                _ = v12
            }
        } else {
            var v13 string = ("hello" + "a")
            delete(v3, (v8 - (^(v8))))
            var v14 int64 = v10
            _ = v13
            _ = v14
        }
    }
    if (!v1) {
        v4 -= v4
        var v15 bool = v1
        fmt.Println(v4, v10, v10)
        _ = v15
    }
    fmt.Println(v1)
    var v16 uint16 = 255
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(fi17, v16)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
