package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }

func main() {
    var v1 string = ("hello" + "")
    var v2 int = 0
    v3 := []uint32{}
    var v4 uint = 255
    v5 := make(map[int]uint32)
    var v6 int16 = 32767
    switch (v2 ^ (v2 + (^(v2)))) {
    case -844:
        if true {
            for fi7 := 0; fi7 < 15; fi7++ {
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 int16 = (v6 | v6)
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 8; fi10++ {
                    _ = len(v3)
                    _ = fi10
                }
                for fi11 := 0; fi11 < 2; fi11++ {
                    if true {
                        var v12 uint16 = 830
                        _ = v12
                    } else {
                        v1 = v1
                        var v13 string = "foo bar"
                        _ = v13
                    }
                    _ = fi11
                }
                _ = fi7
            }
        } else {
            delete(v5, 78)
            v6 = 445
        }
    }
    switch 720 {
    case 32767:
        v5[(v2 | v2)] = 562
    case 948:
        if true {
            v4 = 255
            v14 := GenS1{F0: "hello"}
            _ = v14
        }
    case 128:
        v6 -= v6
    default:
        if false {
            _ = len(v5)
            for fi15 := 0; fi15 < 8; fi15++ {
                if true {
                    var v16 uint32 = 781
                    v1 = (v1 + ("a" + v1))
                    v17 := GenS2{F0: true}
                    _ = v16
                    _ = v17
                }
                _ = len(v3)
                _ = fi15
            }
        }
    }
    v3 = append(v3, 4294967295)
    v18 := GenS3{F0: -852}
    v4 = (v4 ^ 18446744073709551615)
    var v19 int32 = 607
    if true {
        v6 += v6
        _ = len(v5)
    } else {
        switch v1 {
        case "a":
            v2 = v2
        case "hello":
            v1 = "\n\t"
        case "\n\t":
            for fi20 := 0; fi20 < 15; fi20++ {
                switch fi20 {
                case -629:
                    delete(v5, (fi20 & fi20))
                }
                _ = fi20
            }
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int64 = -448
        for fi23 := 0; fi23 < 10; fi23++ {
            if true {
                fmt.Println(v6, v4)
                var v24 bool = true
                var v25 int8 = 127
                _ = v24
                _ = v25
            }
            v6 += v6
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v1 = v1
    v6 = v6
    v2 -= -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
