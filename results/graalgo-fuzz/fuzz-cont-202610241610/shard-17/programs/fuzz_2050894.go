package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int16; F1 int; F2 int64 }
type GenS7 struct { F0 int64 }

func main() {
    v1 := []int{127}
    v2 := make(map[int]bool)
    var v3 uint32 = 0
    var v4 uint = 1
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "a"
        v7 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: 748}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if (18446744073709551615 > uint64(v4)) {
        var v8 int64 = 42
        _ = v8
    } else {
        fmt.Println(v3, v4)
        v9 := GenS2{F0: -2147483648}
        v3 = (v3 - 911)
        _ = v9
    }
    switch ("\n\t" + "\n\t") {
    case "a":
        var v10 string = "a"
        _ = v10
    }
    v11 := GenS3{F0: "foo bar", F1: 235}
    switch 32768 {
    case 247:
        if false {
            if (true && (!false)) {
                v12 := GenS2{F0: -986}
                if (2 == 127) {
                    v13 := GenS2{F0: -9223372036854775808}
                    var v14 uint32 = (v3 & v3)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(v3, v4)
                }
                _ = v12
            } else {
                var v15 uint = 103
                var v16 float64 = 0.1
                _ = v15
                _ = v16
            }
            v4 = 65535
            if (!true) {
                v17 := GenS2{F0: -745}
                var v18 int = 0
                _ = v17
                _ = v18
            }
        }
    }
    var v19 int16 = 1
    v20 := GenS4{F0: 1}
    switch 287 {
    case 42:
        v3 = v3
    default:
        for fi21 := 0; fi21 < 5; fi21++ {
            v4 += v4
            for fi22 := 0; fi22 < 10; fi22++ {
                v23 := GenS5{F0: "hello", F1: 65535}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v24 uint32 = 1
                _ = fi22
                _ = v23
                _ = v24
            }
            v4 += 368
            _ = fi21
        }
    }
    switch "hello" {
    case "\n\t":
        v19 -= (v19 << ((18446744073709551615) & 63))
    }
    _ = len(v1)
    v4 -= (v4 % (((^(v4))) | 1))
    if true {
        for fi25 := 0; fi25 < 3; fi25++ {
            for fi26 := 0; fi26 < 5; fi26++ {
                v4 += 18446744073709551615
                _ = fi26
            }
            var v27 int16 = (v19 & (v19 / ((-162) | 1)))
            _ = fi25
            _ = v27
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            if true {
                fmt.Println(fi28, v3)
                v29 := GenS6{F0: v19, F1: fi28, F2: 9223372036854775807}
                var v30 int32 = 2147483647
                _ = v29
                _ = v30
            } else {
                switch (fi28 + fi28) {
                case 0:
                    switch (fi28 ^ fi28) {
                    case -1:
                        switch "hello" {
                        case "\n\t":
                            v4 = (v4 * v4)
                        case "hello":
                            if (true && (true || true)) {
                                if false {
                                    var v31 uint = 255
                                    _ = v31
                                }
                                var v32 int16 = v19
                                _ = v32
                            }
                        default:
                            v33 := GenS7{F0: 1}
                            _ = v33
                        }
                    }
                case 128:
                    fmt.Println(v19, v4, v3)
                }
            }
            _ = fi28
        }
    }
    var v34 uint64 = 4294967295
    var v35 uint16 = 0
    var v36 int32 = -109
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v37 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
