package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint; F2 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{94, 2147483648}
    var v3 string = "foo bar"
    v4 := []string{}
    var v5 float64 = 0.1
    switch ((v3 + ((v3 + "") + v3)) + v3) {
    case "hello":
        if (!false) {
            v6 := GenS1{F0: false}
            v3 = (((("" + v3) + v3) + v3) + "")
            switch v3 {
            case "\n\t":
                v7 := GenS1{F0: (0 == 1)}
                _ = v7
            }
            _ = v6
        } else {
            var v8 uint64 = 237
            _ = v8
        }
    default:
        v3 = ""
    }
    fmt.Println(v3, v5)
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 uint = 255
            v3 = (v3 + "hello")
            v12 := GenS1{F0: (!true)}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    if true {
        v13 := GenS2{F0: (42 > -120), F1: 2, F2: 597}
        _ = v13
    }
    switch v3 {
    case "foo bar":
        fmt.Println(v5, v5)
    default:
        fmt.Println(v3, v3)
    }
    var v14 string = "foo bar"
    var v15 string = v3
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v15, v3, fi16)
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 int8 = -66
            v14 = ""
            switch 1 {
            case 0:
                var v19 int = (fi16 >> ((18446744073709551615) & 63))
                _ = v19
            case 32768:
                var v20 int16 = 1000
                _ = v20
            case 42:
                v21 := GenS3{F0: 128}
                _ = v21
            }
            v22 := GenS3{F0: 2}
            switch (("foo bar" + v3) + v3) {
            case "":
                v5 += ((float64(1.5) - 0.1) + v5)
            case "\n\t":
                switch (v14 + "hello") {
                case "":
                    fmt.Println(fi17, fi16, v14)
                case "foo bar":
                    var v23 int = 127
                    _ = v23
                default:
                    v18 += v18
                }
            case "a":
                var v24 float32 = 3.14159
                _ = v24
            }
            _ = fi17
            _ = v18
            _ = v22
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
