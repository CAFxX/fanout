package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (("\n\t" + "") != "\n\t")
    v3 := make(map[uint32]string)
    fmt.Println(v2)
    var v4 uint = 42
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint8 = 127
            v7 += v7
            var v8 int32 = -1
            var v9 uint16 = 255
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    if (v2 && (float32(1.5) <= (((1e10 + 1.5) - (-2.25 - -2.25)) - 3.14159))) {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 string = "foo bar"
            if true {
                v11 = v11
                v12 := []int{(fi10 * fi10)}
                fmt.Println(v11, fi10)
                _ = v12
            }
            delete(v1, fi10)
            _ = fi10
            _ = v11
        }
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if (-32768 == 1000) {
            v14 := v1[fi13]
            switch "" {
            case "\n\t":
                var v15 int8 = -65
                _ = v15
            case "hello":
                if false {
                    switch "a" {
                    case "":
                        switch v14 {
                        case 127:
                            v16 := GenS1{F0: 255, F1: 255}
                            _ = v16
                        case 483:
                            fmt.Println(v14)
                        default:
                            v17 := []int{(fi13 + v14)}
                            _ = v17
                        }
                    case "hello":
                        v14 = (fi13 - 42)
                    default:
                        for fi18 := 0; fi18 < 16; fi18++ {
                            v2 = (2 > 4294967295)
                            _ = fi18
                        }
                    }
                    v2 = false
                }
            case "foo bar":
                for fi19 := 0; fi19 < 18; fi19++ {
                    v20 := GenS1{F0: 42, F1: 1}
                    v14 -= fi19
                    fmt.Println(fi13, v2, v2)
                    _ = fi19
                    _ = v20
                }
            default:
                v14 -= -862
            }
            var v21 float64 = ((-2.25 * 1.5) - 3.14159)
            _ = v14
            _ = v21
        } else {
            if (((!(42 < -660)) && v2) && v2) {
                var v22 float64 = 1.5
                _ = v22
            }
            v23 := GenS1{F0: 0, F1: 2}
            v24 := GenS2{F0: -30, F1: fi13, F2: int32(v4)}
            _ = v23
            _ = v24
        }
        v2 = (v2 && ("\n\t" <= ("a" + (("a" + "a") + "\n\t"))))
        var v25 uint32 = 540
        _ = fi13
        _ = v25
    }
    v4 = v4
    if (2 != 0) {
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := []int{}
            v4 += v4
            _ = fi26
            _ = v27
        }
    } else {
        v3[765] = "\n\t"
        v4 = (^(v4))
        v2 = (-2.25 == ((0.1 - 1e10) + 1e10))
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        switch "foo bar" {
        case "\n\t":
            switch ("hello" + "a") {
            case "":
                v29 := GenS3{F0: 42}
                _ = v29
            }
        case "hello":
            for fi30 := 0; fi30 < 16; fi30++ {
                switch "a" {
                case "hello":
                    v31 := GenS3{F0: 65535}
                    _ = v31
                case "foo bar":
                    var v32 int32 = -1
                    _ = v32
                case "":
                    var v33 uint = 65535
                    _ = v33
                }
                fmt.Println(fi30, v4)
                v34 := []int{(fi30 % ((fi30) | 1)), (fi30 >> ((v4) & 63)), (fi28 ^ fi28)}
                _ = fi30
                _ = v34
            }
        case "a":
            if v2 {
                fmt.Println(v2)
            }
        }
        v4 = 2
        var v35 float64 = 1e10
        _ = fi28
        _ = v35
    }
    v36 := []int{-734}
    fmt.Println(v4, v4)
    var v37 string = "hello"
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 int8 = -102
        var v40 uint64 = 147
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
