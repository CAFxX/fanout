package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 string, p1 int16) int {
    v4 := GenS1{F0: p0}
    switch (("" + p0) + p0) {
    case "a":
        for fi5 := 0; fi5 < 4; fi5++ {
            switch "\n\t" {
            case "":
                var v6 uint = 1
                _ = v6
            case "\n\t":
                var v7 string = ""
                _ = v7
            case "hello":
                switch p0 {
                case "hello":
                    if true {
                        var v8 int16 = 32767
                        var v9 int8 = 47
                        _ = v8
                        _ = v9
                    } else {
                        var v10 uint32 = 1
                        _ = v10
                    }
                case "\n\t":
                    for fi11 := 0; fi11 < 10; fi11++ {
                        p1 += (p1 << ((42) & 63))
                        _ = fi11
                    }
                case "foo bar":
                    switch (fi5 ^ 1) {
                    case -194:
                        fmt.Println(p0, p0, p0)
                    }
                }
            default:
                v12 := []int{4294967295, fi5}
                _ = v12
            }
            switch "hello" {
            case "hello":
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS1{F0: "a"}
                    var v15 uint8 = 255
                    switch "" {
                    case "hello":
                        v16 := []int{fi5, (fi13 % (((^(fi5))) | 1))}
                        _ = v16
                    case "foo bar":
                        v17 := []int{}
                        _ = v17
                    case "\n\t":
                        if false {
                            switch fi5 {
                            case -373:
                                var v18 bool = true
                                _ = v18
                            case 4294967295:
                                v19 := []int{-2147483648, (fi5 ^ fi5), (fi13 % (((fi13 & (-(fi13)))) | 1)), fi5}
                                _ = v19
                            }
                            fmt.Println(p0, fi13, fi13)
                            v20 := []int{(fi13 << ((2) & 63)), fi5, fi5}
                            _ = v20
                        }
                    default:
                        fmt.Println(fi13)
                        if true {
                            switch (-(fi13)) {
                            case 300:
                                fmt.Println(fi13)
                            }
                        }
                    }
                    var v21 int8 = 127
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v21
                }
            case "foo bar":
                p0 = "foo bar"
            case "a":
                var v22 int8 = 107
                _ = v22
            }
            _ = fi5
        }
    case "foo bar":
        switch "foo bar" {
        case "":
            var v23 int16 = p1
            _ = v23
        default:
            var v24 float64 = float64(((float32(1.5) - -2.25) * 3.14159))
            _ = v24
        }
    case "\n\t":
        p0 = p0
    default:
        p0 = ""
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}
func genFn1(p0 uint, p1 uint, p2 uint64) int32 {
    if (234 > uint8(p0)) {
        p2 -= p2
    }
    switch int(p0) {
    case -706:
        v25 := GenS1{F0: ("hello" + "\n\t")}
        _ = v25
    case 127:
        v26 := GenS1{F0: ""}
        _ = v26
    default:
        v27 := GenS2{F0: 416, F1: 42}
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    v3 := []uint32{4294967295, 42}
    var v28 string = ("foo bar" + "a")
    var v29 float64 = 0.0
    var v30 int16 = 1000
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint = 0
        fmt.Println(v29)
        _ = fi31
        _ = v32
    }
    v28 = v28
    for fi33 := 0; fi33 < 8; fi33++ {
        fmt.Println(v29, fi33)
        _ = fi33
    }
    if false {
        var v34 int8 = -72
        var v35 int16 = v30
        _ = v34
        _ = v35
    }
    var v36 int8 = -70
    var v37 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
}
