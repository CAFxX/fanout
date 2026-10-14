package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int8; F2 int16 }

func main() {
    var v1 uint16 = 1
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    v4 := []int{}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi5 := 0; fi5 < 15; fi5++ {
        v1 += (^(v1))
        v1 -= v1
        _ = fi5
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 559, F1: 42}
        v1 -= v1
        switch -278 {
        case 32768:
            v8 := GenS1{F0: 595, F1: 683}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    if (42 == 255) {
        v1 = 1
        if false {
            fmt.Println(v1)
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(v1, fi10)
                switch "\n\t" {
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch (("hello" + "foo bar") + "a") {
                    case "":
                        v11 := GenS2{F0: 452}
                        _ = v11
                    case "\n\t":
                        if (!(false && false)) {
                            switch (fi9 >> ((1) & 63)) {
                            case 738:
                                var v12 string = ""
                                _ = v12
                            case 127:
                                if (!false) {
                                    var v13 uint8 = 0
                                    _ = v13
                                }
                            }
                        }
                    }
                case "a":
                    fmt.Println(fi9, fi10)
                case "\n\t":
                    var v14 int64 = 9223372036854775807
                    _ = v14
                }
                _ = fi10
            }
            v15 := GenS1{F0: 497, F1: -9223372036854775808}
            v1 = 2
            _ = fi9
            _ = v15
        }
    } else {
        if (-2147483648 < 1) {
            var v16 uint8 = 2
            _ = len(v4)
            _ = v16
        }
    }
    v17 := GenS3{F0: 26, F1: 42, F2: 193}
    fmt.Println(v1)
    v18 := GenS3{F0: 42, F1: 42, F2: 635}
    v1 = (v1 & v1)
    if false {
        fmt.Println(v1)
        _ = len(v3)
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            v1 = (v1 >> ((1) & 63))
            _ = fi19
        }
    }
    v20 := v3[2]
    v20 = false
    _ = len(v3)
    if v20 {
        if len(v2) > 0 {
            v2[0] = 2
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            v20 = true
            var v22 int = fi21
            _ = fi21
            _ = v22
        }
        if len(v2) > 0 {
            v2[0] = 67
        }
    }
    fmt.Println(v1, v20)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
