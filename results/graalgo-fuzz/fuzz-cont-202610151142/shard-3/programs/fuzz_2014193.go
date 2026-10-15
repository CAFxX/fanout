package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint }
func genFn0() int16 {
    var v7 int16 = -1
    switch "foo bar" {
    case "\n\t":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int32 = 42
            _ = fi8
            _ = v9
        }
    case "a":
        v7 = (^(v7))
    }
    _ = v7
    return -338
}

func main() {
    v1 := []string{("hello" + ""), "", ""}
    v2 := make(map[int]bool)
    var v3 string = "a"
    v4 := []int{737, -981, 0, -719}
    var v5 string = "hello"
    v6 := make(map[uint32]uint32)
    v10 := GenS1{F0: 512, F1: -1}
    fmt.Println(v3, v5)
    for fi11 := 0; fi11 < 3; fi11++ {
        if false {
            for fi12 := 0; fi12 < 10; fi12++ {
                if false {
                    var v13 int8 = 42
                    var v14 uint64 = 0
                    _ = v13
                    _ = v14
                } else {
                    v3 = ""
                    v15 := GenS1{F0: 534, F1: (fi12 & 9223372036854775807)}
                    _ = v15
                }
                v3 = (v5 + v5)
                switch "hello" {
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    var v16 string = "foo bar"
                    _ = v16
                }
                _ = fi12
            }
            if false {
                var v17 int64 = -69
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            switch (v5 + v5) {
            case "hello":
                var v19 uint64 = 753
                _ = v19
            case "a":
                v3 = ""
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    if true {
                        if false {
                            var v20 int64 = int64(fi18)
                            _ = v20
                        }
                    } else {
                        fmt.Println(fi18, fi18)
                        v3 = v3
                    }
                case "hello":
                    v21 := v6[57]
                    _ = v21
                default:
                    switch ("\n\t" + ("\n\t" + (v5 + ("a" + v5)))) {
                    case "a":
                        v22 := GenS2{F0: 1, F1: 0}
                        _ = v22
                    case "hello":
                        v23 := GenS3{F0: 42, F1: 203}
                        _ = v23
                    }
                }
            default:
                var v24 uint = 2
                fmt.Println(v5, fi18)
                var v25 uint16 = 32767
                _ = v24
                _ = v25
            }
            _ = fi18
        }
        v26 := GenS4{F0: -128, F1: 2}
        _ = fi11
        _ = v26
    }
    var v27 int16 = -15
    v5 = v3
    for fi28 := 0; fi28 < 20; fi28++ {
        if ((1 != 128) && (!true)) {
            var v29 int8 = 127
            v4 = append(v4, (-(fi28)))
            _ = v29
        } else {
            var v30 uint = 41
            fmt.Println(v27, v30)
            _ = v30
        }
        delete(v6, 2147483647)
        v5 = v3
        _ = fi28
    }
    var v31 int64 = 871
    v31 -= 136
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
