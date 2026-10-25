package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint16) string {
    var v5 uint64 = (p0 & (^(p0)))
    var v6 uint8 = 219
    v7 := []int{1}
    v6 += (v6 >> ((42) & 63))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{1, 65535, 2147483647, 2147483647}
    v3 := []uint32{42, 426, 65535}
    v4 := []string{"foo bar"}
    switch -1 {
    case -601:
        switch 1 {
        case 0:
            var v8 int = 1
            _ = v8
        }
    case -1:
        v9 := GenS1{F0: 32767}
        _ = v9
    }
    var v10 uint16 = 32767
    v11 := GenS1{F0: -401}
    v10 = 133
    var v12 int = 32767
    fmt.Println(v10, v12)
    for fi13 := 0; fi13 < 1; fi13++ {
        if (255 <= 1) {
            if false {
                fmt.Println(v10, v12)
                v14 := GenS1{F0: 42}
                if true {
                    var v15 int32 = -123
                    if true {
                        v16 := GenS2{F0: -32768, F1: 4294967295, F2: 0}
                        _ = v16
                    }
                    if ((float64(1e10) == (-2.25 - 1.5)) || true) {
                        var v17 uint16 = (v10 / (((v10 * v10)) | 1))
                        var v18 uint8 = 85
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                }
                _ = v14
            }
            v19 := GenS3{F0: 42, F1: 763, F2: 32767}
            switch "a" {
            case "a":
                if false {
                    var v20 uint32 = 1
                    _ = v20
                } else {
                    v1[2147483648] = ("hello" + "foo bar")
                }
            default:
                for fi21 := 0; fi21 < 8; fi21++ {
                    fmt.Println(fi13, v12, fi21)
                    _ = fi21
                }
            }
            _ = v19
        } else {
            v22 := GenS1{F0: 42}
            if len(v2) > 0 {
                v2[0] = 65535
            }
            _ = v22
        }
        v12 = fi13
        _ = fi13
    }
    switch ("hello" + "hello") {
    case "":
        for fi23 := 0; fi23 < 13; fi23++ {
            switch (("a" + "\n\t") + "\n\t") {
            case "\n\t":
                fmt.Println(fi23, v12)
            case "foo bar":
                for fi24 := 0; fi24 < 14; fi24++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v25 float32 = float32(1.5)
                    _ = fi24
                    _ = v25
                }
            }
            _ = fi23
        }
    case "\n\t":
        for fi26 := 0; fi26 < 4; fi26++ {
            if (2 <= 4294967295) {
                switch "foo bar" {
                case "foo bar":
                    var v27 int = v12
                    _ = v27
                default:
                    v10 -= (v10 * 1)
                }
                switch "hello" {
                case "hello":
                    if (!false) {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            if true {
                                v29 := GenS4{F0: ""}
                                _ = v29
                            }
                            _ = fi28
                        }
                        _ = len(v1)
                    } else {
                        var v30 uint8 = 2
                        _ = v30
                    }
                case "":
                    fmt.Println(v10, fi26)
                default:
                    var v31 uint32 = 26
                    _ = v31
                }
            } else {
                for fi32 := 0; fi32 < 12; fi32++ {
                    if (!(100 > 1)) {
                        if (v10 < v10) {
                            v12 = fi32
                            v10 += (v10 >> ((0) & 63))
                            v10 += 1
                        }
                    } else {
                        var v33 int16 = 256
                        var v34 uint = 618
                        v12 -= 32767
                        _ = v33
                        _ = v34
                    }
                    v12 = v12
                    _ = fi32
                }
                v35 := GenS5{F0: true}
                for fi36 := 0; fi36 < 13; fi36++ {
                    if true {
                        var v37 uint = uint(v10)
                        _ = v37
                    } else {
                        v38 := GenS3{F0: -1, F1: 780, F2: 2}
                        fmt.Println(v12)
                        v39 := GenS6{F0: 2147483647, F1: -263}
                        _ = v38
                        _ = v39
                    }
                    _ = fi36
                }
                _ = v35
            }
            if (!(0 <= v12)) {
                var v40 string = "foo bar"
                v41 := GenS7{F0: 560}
                fmt.Println(fi26, v40, fi26)
                _ = v40
                _ = v41
            }
            _ = fi26
        }
    case "foo bar":
        switch "a" {
        case "hello":
            var v42 float64 = 1.5
            _ = v42
        }
    }
    if false {
        var v43 int = (v12 | v12)
        v12 += (^(v43))
        _ = v43
    } else {
        v12 = (v12 | v12)
        var v44 uint = 430
        v1[2147483647] = "foo bar"
        _ = v44
    }
    var v45 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
