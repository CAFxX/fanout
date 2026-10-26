package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint32, p1 int, p2 uint64) uint64 {
    switch ("foo bar" + ("hello" + "hello")) {
    case "a":
        p2 -= uint64(p0)
    case "":
        var v7 uint64 = 2
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        if false {
            if false {
                var v9 int = 1
                v10 := GenS1{F0: 2, F1: 0}
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 uint16 = 32767
                    _ = fi11
                    _ = v12
                }
                var v13 uint32 = p0
                v14 := GenS1{F0: 42, F1: -638}
                _ = v13
                _ = v14
            }
            switch (fi8 % (((fi8 & (p1 % (((p1 - 68)) | 1)))) | 1)) {
            case 237:
                for fi15 := 0; fi15 < 9; fi15++ {
                    p0 -= p0
                    _ = fi15
                }
            }
        }
        var v16 uint32 = p0
        _ = fi8
        _ = v16
    }
    v17 := GenS2{F0: (("a" + "hello") + "\n\t")}
    p0 += p0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (^(p2))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (42 > 42)
    v3 := []string{}
    v4 := make(map[int]bool)
    var v5 int16 = -1
    var v6 uint8 = 128
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint16 = 248
        delete(v4, fi18)
        _ = fi18
        _ = v19
    }
    var v20 int8 = 100
    v6 = (v6 >> ((255) & 63))
    if v2 {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 uint32 = 465
            v22 -= (^(v22))
            v23 := GenS3{F0: 1, F1: 0}
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 int64 = -458
        if v2 {
            for fi25 := 0; fi25 < 11; fi25++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch (("hello" + (("\n\t" + ("" + "hello")) + "\n\t")) + "hello") {
                case "a":
                    if (!true) {
                        switch "a" {
                        case "\n\t":
                            var v26 int32 = -1
                            _ = v26
                        }
                        var v27 float32 = 3.14159
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 13; fi28++ {
                            if false {
                                _ = len(v3)
                                var v29 int64 = v24
                                var v30 uint8 = 137
                                _ = v29
                                _ = v30
                            }
                            _ = fi28
                        }
                    }
                }
                _ = fi25
            }
        }
        _ = v24
    }
    v6 -= (v6 + v6)
    v2 = v2
    var v31 uint8 = 255
    fmt.Println(v20, v31)
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 int8 = (v20 & 1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
