package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int64, p2 uint64) string {
    p2 += (p0 | (p2 + (p0 | 362)))
    var v7 uint = 972
    v7 -= v7
    var v8 uint16 = 42
    var v9 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ("\n\t" + "hello")
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint32 = 2
    var v3 int32 = -1
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    v6 := []uint32{}
    switch 505 {
    case 128:
        fmt.Println(v3, v2)
    case 32767:
        var v10 string = ""
        _ = v10
    case -9223372036854775808:
        v3 = v3
    default:
        var v11 int8 = -1
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        if true {
            var v13 int16 = 1000
            v2 = v2
            switch "\n\t" {
            case "a":
                if false {
                    var v14 int = (fi12 / ((-38) | 1))
                    v15 := GenS1{F0: 42}
                    fmt.Println(v3, v14)
                    _ = v14
                    _ = v15
                } else {
                    var v16 int64 = 0
                    fmt.Println(v13, v16, fi12)
                    _ = v16
                }
            case "\n\t":
                var v17 int8 = 42
                _ = v17
            case "foo bar":
                var v18 int64 = -9223372036854775808
                _ = v18
            default:
                v1 = append(v1, ((("foo bar" + ("hello" + "\n\t")) + "hello") + "\n\t"))
            }
            _ = v13
        } else {
            v4[(v2 % ((65535) | 1))] = true
        }
        v19 := GenS2{F0: 193}
        fmt.Println(fi12)
        _ = fi12
        _ = v19
    }
    fmt.Println(v3, v3)
    var v20 int32 = 2147483647
    if (!true) {
        for fi21 := 0; fi21 < 20; fi21++ {
            fmt.Println(v20)
            _ = fi21
        }
        if len(v1) > 0 {
            v1[0] = (("a" + "foo bar") + ((("" + (("hello" + "a") + "a")) + "") + "foo bar"))
        }
        if false {
            var v22 uint8 = 1
            v3 = (v20 ^ (v3 + v3))
            _ = v22
        } else {
            v20 -= 42
        }
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v2 = v2
        _ = fi23
    }
    fmt.Println(v3, v3, v2)
    if ("" < "hello") {
        for fi24 := 0; fi24 < 9; fi24++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            v25 := GenS2{F0: 0}
            var v26 int = (fi24 | fi24)
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 int16 = 42
        delete(v5, ("\n\t" + "\n\t"))
        _ = v27
    }
    v3 += (v3 >> ((0) & 63))
    switch "hello" {
    case "hello":
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 uint32 = 2
            if (0.1 != 3.14159) {
                var v30 float64 = 3.14159
                _ = v30
            } else {
                for fi31 := 0; fi31 < 16; fi31++ {
                    v32 := GenS2{F0: 255}
                    var v33 uint8 = 0
                    var v34 uint32 = 753
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                if true {
                    switch "" {
                    case "hello":
                        v5["hello"] = fi28
                    case "foo bar":
                        v35 := GenS2{F0: 2}
                        _ = v35
                    default:
                        v2 -= v29
                    }
                    _ = len(v5)
                    var v36 bool = true
                    _ = v36
                }
            }
            _ = fi28
            _ = v29
        }
    case "\n\t":
        switch -1 {
        case -2147483648:
            switch 1 {
            case 4294967295:
                var v37 int8 = -128
                _ = v37
            case 42:
                v4[42] = false
            case 128:
                if (!((825 == 9223372036854775807) && (false || false))) {
                    v38 := GenS2{F0: 0}
                    _ = v38
                } else {
                    var v39 float32 = 0.0
                    _ = v39
                }
            }
        case 9223372036854775807:
            v20 -= 0
        }
    }
    fmt.Println(v20, v20)
    for fi40 := 0; fi40 < 15; fi40++ {
        for fi41 := 0; fi41 < 13; fi41++ {
            var v42 uint16 = 179
            if true {
                v42 = (v42 ^ v42)
            }
            var v43 bool = true
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi40
    }
    delete(v5, "a")
    v3 = v20
    var v44 int = -933
    if false {
        fmt.Println(v2)
    } else {
        var v45 bool = false
        _ = v45
    }
    var v46 int8 = -55
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
    _ = v6
}
