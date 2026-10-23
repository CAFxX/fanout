package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32, p1 int32) uint {
    var v4 uint32 = 65535
    var v5 int8 = 17
    var v6 uint16 = 2
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -938
    var v3 int32 = int32(v2)
    var v7 int = v2
    var v8 uint16 = 42
    for fi9 := 0; fi9 < 10; fi9++ {
        if false {
            v2 = (fi9 << ((65535) & 63))
            switch fi9 {
            case 32768:
                for fi10 := 0; fi10 < 9; fi10++ {
                    switch "" {
                    case "foo bar":
                        switch "" {
                        case "foo bar":
                            if (!(!true)) {
                                v11 := GenS1{F0: 2}
                                fmt.Println(v7, v2, v8)
                                v7 = v2
                                _ = v11
                            }
                        }
                    }
                    fmt.Println(v2, v3)
                    _ = fi10
                }
            case 42:
                var v12 bool = ("hello" == "foo bar")
                _ = v12
            default:
                fmt.Println(fi9, v3, v8)
            }
        }
        _ = fi9
    }
    v13 := GenS1{F0: 18446744073709551615}
    switch ((("foo bar" + ("\n\t" + "")) + (("foo bar" + "\n\t") + "a")) + ("foo bar" + "foo bar")) {
    case "":
        var v14 uint32 = 2147483647
        _ = v14
    default:
        v15 := GenS1{F0: 42}
        _ = v15
    }
    fmt.Println(v8, v2, v8)
    v8 += v8
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint32 = uint32(v8)
            fmt.Println(fi16, fi17, v2)
            _ = fi17
            _ = v18
        }
        var v19 bool = true
        _ = fi16
        _ = v19
    }
    switch "" {
    case "foo bar":
        if false {
            v2 = v7
        }
    case "a":
        fmt.Println(v8, v7, v2)
    }
    fmt.Println(v3, v2, v7)
    var v20 float64 = 3.14159
    v21 := GenS2{F0: 505}
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(v20)
        _ = fi22
    }
    var v23 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v7
    _ = v8
}
