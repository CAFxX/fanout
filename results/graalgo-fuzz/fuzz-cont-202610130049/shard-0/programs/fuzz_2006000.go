package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 string = "foo bar"
    v2 := []int{}
    var v3 uint32 = 631
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi4 := 0; fi4 < 19; fi4++ {
        v1 = "a"
        _ = fi4
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "":
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 uint8 = 128
            v7 := GenS1{F0: -1}
            var v8 int16 = 42
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    default:
        v1 = ""
    }
    fmt.Println(v1, v1)
    if ((!true) && (42 > 18446744073709551615)) {
        v3 += 687
        fmt.Println(v3, v1)
    } else {
        switch "hello" {
        case "a":
            v3 -= 1
        case "hello":
            var v9 uint64 = 0
            _ = v9
        case "foo bar":
            var v10 int8 = 64
            _ = v10
        }
        v11 := GenS2{F0: 1000}
        _ = v11
    }
    v12 := GenS3{F0: 1}
    v3 = v3
    v13 := GenS2{F0: -937}
    var v14 int = 32767
    _ = len(v2)
    var v15 int16 = -1
    v3 += v3
    if true {
        if true {
            var v16 uint32 = 193
            if ((true || (315 == v15)) || false) {
                fmt.Println(v3)
                switch (v14 << ((4294967295) & 63)) {
                case 32768:
                    v16 += v16
                case 128:
                    v17 := GenS1{F0: -32768}
                    _ = v17
                }
                var v18 int32 = -2147483648
                _ = v18
            } else {
                v19 := GenS2{F0: v15}
                v14 = (v14 << ((2) & 63))
                _ = v19
            }
            _ = v16
        }
    }
    var v20 float64 = float64((float32(0.1) * 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}
