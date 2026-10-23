package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0(p0 int8, p1 uint32, p2 int) uint {
    if false {
        switch "\n\t" {
        case "foo bar":
            var v6 uint64 = uint64(p1)
            _ = v6
        case "\n\t":
            fmt.Println(p1, p1, p0)
        case "":
            p0 = (p0 * (^(p0)))
        }
        v7 := []int{(^(p2))}
        var v8 uint = 4294967295
        _ = v7
        _ = v8
    }
    if true {
        p2 += p2
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 string = "hello"
    var v2 int64 = 0
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 uint64 = 42
    if (!((32767 != 127) && (0 != 1))) {
        v5 -= v5
        fmt.Println(v1, v1)
    }
    v2 = 42
    v5 = 4294967295
    v5 = v5
    if false {
        if false {
            v9 := []int{-9223372036854775808, 32767}
            if (-954 < 1) {
                v1 = ((((((v1 + ("foo bar" + ((v1 + "hello") + ""))) + "hello") + v1) + "hello") + "") + v1)
            } else {
                v10 := GenS1{F0: 1, F1: -179}
                delete(v4, (("\n\t" + v1) + (v1 + ("foo bar" + (v1 + v1)))))
                _ = v10
            }
            var v11 uint = 821
            _ = v11
            _ = v9
        } else {
            if true {
                fmt.Println(v2)
                var v12 int32 = 2147483647
                _ = v12
            } else {
                var v13 int16 = 1000
                if (0.0 > float64(0.1)) {
                    var v14 uint16 = 0
                    v15 := GenS1{F0: 42, F1: 42}
                    var v16 int = -1
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
            v5 += (v5 >> ((123) & 63))
        }
        switch 199 {
        case 32768:
            if true {
                v4[("\n\t" + v1)] = 205
                switch v1 {
                case "\n\t":
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v2 += (^(v2))
                        _ = fi17
                    }
                case "hello":
                    switch ((("hello" + v1) + ("\n\t" + v1)) + "a") {
                    case "a":
                        v1 = v1
                    }
                case "foo bar":
                    v18 := []int{-769, 862, 127}
                    _ = v18
                default:
                    var v19 uint8 = 109
                    _ = v19
                }
            } else {
                v5 += (v5 | v5)
            }
        case 128:
            switch "\n\t" {
            case "hello":
                v20 := GenS2{F0: ((v1 != ((v1 + "a") + v1)) && (4294967295 >= 42)), F1: (v5 % ((850) | 1)), F2: 0}
                _ = v20
            }
        case -224:
            var v21 int32 = 264
            _ = v21
        }
    }
    var v22 uint32 = 42
    v5 += v5
    v23 := GenS3{F0: 225, F1: 1}
    var v24 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
