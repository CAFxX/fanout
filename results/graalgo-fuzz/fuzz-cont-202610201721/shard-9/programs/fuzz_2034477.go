package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 int64, p1 bool) string {
    if (true && (p1 || p1)) {
        var v5 bool = (p1 && p1)
        p1 = p1
        _ = v5
    } else {
        p1 = p1
        fmt.Println(p0, p1)
    }
    p0 = p0
    p1 = (2 > 1)
    _ = p0
    _ = p1
    return (("foo bar" + "foo bar") + "foo bar")
}

func main() {
    var v1 uint8 = 166
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 int = 0
    switch "hello" {
    case "foo bar":
        var v6 uint16 = 0
        _ = v6
    case "hello":
        fmt.Println(v4, v1)
    }
    var v7 int8 = 0
    v8 := GenS1{F0: "hello", F1: 1, F2: 42}
    var v9 int16 = -1
    switch "" {
    case "foo bar":
        switch -65 {
        case -1:
            v3[(v4 << ((958) & 63))] = 0
        case 9223372036854775807:
            v3[-9223372036854775808] = 2
        }
    }
    v4 -= (v4 & v4)
    var v10 int16 = (v9 % ((v9) | 1))
    if (!(996 <= 0)) {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint8 = (^(v1))
            if (("foo bar" == ("foo bar" + (("foo bar" + "\n\t") + "a"))) && true) {
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := v2[32767]
                    var v15 uint64 = 18446744073709551615
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: "a", F1: 42, F2: 32767}
                _ = v16
            }
            v17 := GenS2{F0: v10, F1: 4294967295, F2: 498}
            _ = fi11
            _ = v12
            _ = v17
        }
        v18 := []int{2147483647, v4, (v4 >> ((895) & 63)), (v4 - int(v9))}
        if (!(!true)) {
            for fi19 := 0; fi19 < 20; fi19++ {
                v20 := GenS2{F0: (v9 | (v9 << ((480) & 63))), F1: 714, F2: 65535}
                if (!((2147483647 <= 651) || (2147483647 > 1))) {
                    v7 = 1
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v7 = (v7 & (v7 * v7))
                        if (1 >= 95) {
                            v1 -= (^(v1))
                        }
                        _ = fi21
                    }
                    for fi22 := 0; fi22 < 14; fi22++ {
                        switch "hello" {
                        case "foo bar":
                            v7 += 127
                        }
                        _ = fi22
                    }
                }
                _ = fi19
                _ = v20
            }
            if (("" + "a") >= (("" + "hello") + "foo bar")) {
                fmt.Println(v9)
            } else {
                if false {
                    var v23 uint32 = 4294967295
                    _ = v23
                }
                var v24 int32 = -20
                _ = v24
            }
        } else {
            v25 := GenS3{F0: (v7 - -128), F1: 255}
            delete(v2, v4)
            _ = v25
        }
        _ = v18
    }
    fmt.Println(v7, v7, v9)
    for fi26 := 0; fi26 < 8; fi26++ {
        if (!true) {
            v10 = v10
        } else {
            v27 := GenS2{F0: (-(v10)), F1: 4294967295, F2: 550}
            if false {
                if true {
                    var v28 float32 = 1.5
                    v10 = v10
                    _ = v28
                }
                v10 -= 0
            }
            v10 += 42
            _ = v27
        }
        _ = fi26
    }
    switch "foo bar" {
    case "a":
        if (((true || false) || false) && false) {
            fmt.Println(v1, v4)
        }
    case "\n\t":
        for fi29 := 0; fi29 < 9; fi29++ {
            switch "foo bar" {
            case "a":
                for fi30 := 0; fi30 < 9; fi30++ {
                    v7 = 100
                    v31 := GenS3{F0: (v7 << ((2) & 63)), F1: 879}
                    _ = fi30
                    _ = v31
                }
            case "foo bar":
                fmt.Println(v4, v1, v10)
            case "hello":
                v32 := []int{v4, (v4 | fi29), fi29}
                _ = v32
            default:
                var v33 bool = false
                _ = v33
            }
            _ = fi29
        }
    case "foo bar":
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
