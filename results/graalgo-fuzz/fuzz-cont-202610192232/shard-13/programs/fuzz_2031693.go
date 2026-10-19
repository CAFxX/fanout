package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 string, p2 uint16) int {
    var v7 uint32 = 2
    fmt.Println(p2, v7, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}
func genFn1(p0 bool, p1 bool) int64 {
    p0 = (p1 || true)
    fmt.Println(p0, p1)
    var v8 int32 = 2147483647
    p1 = p1
    _ = p0
    _ = p1
    _ = v8
    return -969
}

func main() {
    var v1 uint32 = 176
    var v2 bool = (1.5 == 1e10)
    var v3 uint32 = (^(v1))
    v4 := []string{("a" + (("hello" + "foo bar") + ("hello" + "\n\t"))), "a", ("a" + "\n\t"), "\n\t"}
    v5 := []int{-788, 42, 1}
    var v6 bool = true
    switch (("a" + ("a" + "\n\t")) + "foo bar") {
    case "a":
        switch "a" {
        case "a":
            switch "a" {
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "":
                fmt.Println(v1, v2, v1)
            default:
                switch "" {
                case "\n\t":
                    var v9 float32 = 3.14159
                    _ = v9
                case "a":
                    v5 = append(v5, 391)
                default:
                    var v10 int8 = 1
                    _ = v10
                }
            }
        case "\n\t":
            v1 = (v3 >> ((4294967295) & 63))
        default:
            v4 = append(v4, "\n\t")
        }
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        switch -160 {
        case 4294967295:
            if ((v6 || v2) || false) {
                var v11 bool = v6
                v12 := GenS1{F0: 1}
                _ = v11
                _ = v12
            }
        case 42:
            v5 = append(v5, 9223372036854775807)
        case 701:
            v13 := GenS1{F0: 0}
            _ = v13
        }
    }
    switch 0 {
    case 0:
        if v2 {
            var v14 bool = (((v2 && true) || (!(!v2))) && false)
            _ = v14
        }
    default:
        v5 = append(v5, -428)
    }
    var v15 uint = 65535
    if len(v5) > 0 {
        v5[0] = 404
    }
    if len(v5) > 0 {
        v5[0] = 32767
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v2 = (false && (!(!(false && (4294967295 != 42)))))
        v17 := GenS1{F0: -9223372036854775808}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 1, F1: "\n\t", F2: 42}
    if v2 {
        if v6 {
            var v19 uint32 = 2147483648
            var v20 float32 = 0.1
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 0}
            v22 := GenS2{F0: int(v3), F1: ("\n\t" + "foo bar"), F2: 0}
            fmt.Println(v15, v1, v15)
            _ = v21
            _ = v22
        }
        fmt.Println(v1, v6, v15)
        for fi23 := 0; fi23 < 20; fi23++ {
            v1 -= v3
            v24 := GenS3{F0: 255}
            _ = fi23
            _ = v24
        }
    }
    var v25 uint32 = (v1 & (v3 ^ v1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
