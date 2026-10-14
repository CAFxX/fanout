package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 int32) int8 {
    v5 := GenS1{F0: -1, F1: 4294967295, F2: 65535}
    p0 = p0
    v6 := []int{-2147483648, 2147483647}
    for fi7 := 0; fi7 < 5; fi7++ {
        switch fi7 {
        case -9223372036854775808:
            switch "hello" {
            case "":
                var v8 string = "\n\t"
                _ = v8
            case "foo bar":
                var v9 int64 = 322
                _ = v9
            }
        case 546:
            var v10 int32 = (^(p0))
            _ = v10
        case -1:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        var v11 uint64 = 42
        _ = fi7
        _ = v11
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 int16) int64 {
    p0 = p1
    v12 := []int{9223372036854775807, -989, 2147483647}
    var v13 int64 = -9223372036854775808
    if (127 <= 127) {
        switch -36 {
        case 32767:
            v14 := GenS1{F0: 1, F1: 255, F2: 0}
            _ = v14
        case 2147483647:
            var v15 uint = 1
            _ = v15
        case 1:
            p1 = (p0 + (p1 | (p0 % ((32767) | 1))))
        }
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return (v13 << ((0) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 int64 = -435
    delete(v1, 576)
    var v16 string = "a"
    if true {
        var v17 float32 = 3.14159
        var v18 float64 = 3.14159
        _ = v17
        _ = v18
    }
    switch (v16 + "hello") {
    case "a":
        v19 := GenS1{F0: 42, F1: 4, F2: 221}
        _ = v19
    case "":
        delete(v1, 823)
    }
    switch v16 {
    case "\n\t":
        var v20 uint8 = 255
        _ = v20
    case "a":
        var v21 int32 = -313
        _ = v21
    case "hello":
        fmt.Println(v16)
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 uint32 = 0
            var v25 uint32 = (v24 + v24)
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 int16 = 0
        _ = fi22
        _ = v26
    }
    switch -2147483648 {
    case -2147483648:
        for fi27 := 0; fi27 < 15; fi27++ {
            delete(v1, 4294967295)
            for fi28 := 0; fi28 < 9; fi28++ {
                v16 = (v16 + v16)
                _ = fi28
            }
            var v29 int32 = 2147483647
            _ = fi27
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        v16 = (("\n\t" + v16) + v16)
        _ = fi30
    }
    switch v16 {
    case "hello":
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := v1[0]
            _ = fi31
            _ = v32
        }
    case "":
        v16 = ("\n\t" + v16)
    case "foo bar":
        for fi33 := 0; fi33 < 8; fi33++ {
            fmt.Println(fi33, v16)
            _ = fi33
        }
    }
    v34 := v1[2147483647]
    if false {
        var v35 float32 = (0.0 - 1.5)
        _ = v35
    } else {
        switch (v16 + ("hello" + "\n\t")) {
        case "":
            var v36 int = (v34 & v34)
            _ = v36
        case "\n\t":
            for fi37 := 0; fi37 < 10; fi37++ {
                delete(v2, 0)
                _ = fi37
            }
        case "a":
            v38 := GenS2{F0: 1, F1: 1, F2: 1}
            _ = v38
        default:
            v39 := []int{(v34 % ((127) | 1)), 32767, 655, v34}
            _ = v39
        }
    }
    v40 := []int{-616, -231, (v34 << ((1) & 63)), -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
}
