package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -117
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    v2 += 42
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 bool = (26 >= -128)
        switch fi5 {
        case 127:
            v7 := []int{fi5, (fi5 + (fi5 << ((65535) & 63)))}
            _ = v7
        case -101:
            if (!v6) {
                fmt.Println(v2, v6)
            } else {
                for fi8 := 0; fi8 < 13; fi8++ {
                    var v9 string = "foo bar"
                    _ = fi8
                    _ = v9
                }
            }
        }
        _ = fi5
        _ = v6
    }
    if true {
        v10 := GenS1{F0: 1}
        var v11 int64 = 42
        _ = v10
        _ = v11
    }
    if (0 <= -366) {
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v2, fi12)
            fmt.Println(fi12, v2)
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        if (!true) {
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS1{F0: -9223372036854775808}
                fmt.Println(v2)
                _ = fi14
                _ = v15
            }
            if (0 == 255) {
                v16 := GenS2{F0: 18446744073709551615}
                _ = v16
            } else {
                var v17 int64 = 42
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 uint = 262
                    v2 = (v2 | -2147483648)
                    var v20 bool = false
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v2)
                _ = v17
            }
        }
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v21 := GenS2{F0: 0}
    var v22 uint16 = 144
    if (true && (true && false)) {
        fmt.Println(v22, v22)
        _ = len(v3)
    } else {
        v23 := []int{-9223372036854775808}
        _ = v23
    }
    if true {
        fmt.Println(v22)
        switch ("a" + "") {
        case "a":
            v24 := GenS3{F0: 37, F1: "a", F2: 4294967295}
            _ = v24
        case "hello":
            v25 := GenS4{F0: 128}
            _ = v25
        }
    }
    fmt.Println(v2, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
