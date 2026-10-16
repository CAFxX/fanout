package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
type GenS4 struct { F0 int; F1 int32; F2 uint }

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]int)
    v3 := []int{9223372036854775807, -9223372036854775808, 42, -2147483648}
    var v4 int16 = -969
    v5 := []string{}
    v6 := make(map[int]string)
    switch -1 {
    case 9223372036854775807:
        if (32767 < 65535) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 bool = true
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "hello" {
                case "":
                    fmt.Println(v8, v4, fi7)
                case "hello":
                    _ = len(v5)
                case "foo bar":
                    var v9 int = 1
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 12; fi10++ {
                        if true {
                            v11 := GenS1{F0: 0, F1: 42, F2: 886}
                            _ = v11
                        }
                        var v12 int8 = 1
                        _ = fi10
                        _ = v12
                    }
                }
                _ = fi7
                _ = v8
            }
        }
    case 128:
        switch -9223372036854775808 {
        case -2147483648:
            var v13 uint16 = 255
            _ = v13
        }
    case 2147483647:
        v14 := GenS2{F0: 1}
        _ = v14
    }
    _ = len(v5)
    var v15 uint = v1
    if false {
        v16 := GenS3{F0: "hello", F1: 1, F2: 4294967295}
        fmt.Println(v15, v15, v1)
        if true {
            v17 := GenS4{F0: 32767, F1: 42, F2: 18446744073709551615}
            for fi18 := 0; fi18 < 12; fi18++ {
                _ = len(v6)
                var v19 float64 = (1.5 - 1.5)
                _ = fi18
                _ = v19
            }
            _ = v17
        } else {
            var v20 string = "hello"
            for fi21 := 0; fi21 < 8; fi21++ {
                fmt.Println(fi21)
                var v22 uint64 = 2
                _ = fi21
                _ = v22
            }
            switch (v20 + (v20 + "foo bar")) {
            case "\n\t":
                var v23 int16 = v4
                _ = v23
            case "a":
                v4 = v4
            case "":
                _ = len(v6)
            }
            _ = v20
        }
        _ = v16
    }
    var v24 uint16 = 32767
    if (!(!false)) {
        v25 := GenS2{F0: (-(v4))}
        v3 = append(v3, 1)
        _ = v25
    }
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    fmt.Println(v15, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
