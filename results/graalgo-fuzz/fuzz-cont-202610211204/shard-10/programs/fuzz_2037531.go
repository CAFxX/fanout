package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int16) int {
    var v6 float32 = 0.0
    var v7 uint = 65535
    v6 += v6
    var v8 uint64 = 407
    v8 -= v8
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int) int16 {
    fmt.Println(p0)
    p0 = (p0 / ((9223372036854775807) | 1))
    fmt.Println(p0)
    _ = p0
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2147483647, 148}
    v3 := make(map[int]string)
    var v4 uint = 18446744073709551615
    var v5 bool = (!(0 < 2))
    switch "hello" {
    case "a":
        fmt.Println(v4, v4)
    }
    var v9 string = (("hello" + "\n\t") + ("a" + "a"))
    if (!v5) {
        for fi10 := 0; fi10 < 12; fi10++ {
            v5 = (0 < -32768)
            for fi11 := 0; fi11 < 16; fi11++ {
                _ = len(v3)
                v4 = (^(v4))
                switch (^(fi10)) {
                case 887:
                    v12 := GenS1{F0: 42}
                    switch v9 {
                    case "foo bar":
                        fmt.Println(v5)
                    }
                    delete(v1, v9)
                    _ = v12
                case -2147483648:
                    v13 := v1["a"]
                    _ = v13
                default:
                    v9 = (v9 + (v9 + "foo bar"))
                }
                _ = fi11
            }
            v3[(fi10 % ((32768) | 1))] = "\n\t"
            _ = fi10
        }
    } else {
        var v14 float32 = 1e10
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if v5 {
            var v16 uint16 = 65535
            for fi17 := 0; fi17 < 12; fi17++ {
                if false {
                    var v18 uint8 = 42
                    if (3.14159 == 0.1) {
                        fmt.Println(v9, fi17, v18)
                    }
                    _ = v18
                } else {
                    _ = len(v2)
                }
                var v19 int64 = 446
                v20 := GenS1{F0: 42}
                switch ("\n\t" + v9) {
                case "foo bar":
                    v9 = ("hello" + v9)
                case "a":
                    v21 := GenS2{F0: 18446744073709551615}
                    _ = v21
                case "hello":
                    v16 = 255
                }
                _ = fi17
                _ = v19
                _ = v20
            }
            _ = v16
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := GenS2{F0: 287}
            v24 := GenS2{F0: 4294967295}
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi15
    }
    switch 120 {
    case 2147483647:
        for fi25 := 0; fi25 < 18; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                delete(v3, genFn0(1, -1))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi26
            }
            fmt.Println(fi25, v9)
            v27 := GenS1{F0: 4}
            _ = fi25
            _ = v27
        }
    case 127:
        fmt.Println(v5)
    case -9223372036854775808:
        fmt.Println(v4, v5, v5)
    }
    v28 := GenS2{F0: 42}
    switch 9223372036854775807 {
    case -2147483648:
        switch v9 {
        case "\n\t":
            var v29 uint64 = 42
            _ = v29
        case "a":
            var v30 int = 128
            _ = v30
        case "":
            v5 = (((v4 << (((v4 ^ (v4 >> (((v4 >> ((42) & 63))) & 63)))) & 63)) >= v4) || v5)
        default:
            if v5 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v31 := GenS1{F0: 255}
                _ = v31
            }
        }
    case -953:
        v9 = ("hello" + v9)
    }
    delete(v1, v9)
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
