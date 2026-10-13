package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int64 }
func genFn0() uint16 {
    var v4 uint8 = 2
    fmt.Println(v4)
    v5 := GenS1{F0: 42, F1: 255}
    v4 = 2
    v6 := GenS1{F0: 839, F1: uint(v4)}
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 string = ("hello" + ("a" + "a"))
    for fi7 := 0; fi7 < 16; fi7++ {
        if true {
            fmt.Println(fi7, v3)
            v8 := []int{fi7, 32767, (fi7 & fi7)}
            _ = v8
        }
        if (18446744073709551615 > 1) {
            for fi9 := 0; fi9 < 20; fi9++ {
                delete(v1, fi9)
                v3 = "foo bar"
                _ = fi9
            }
            if true {
                v3 = v3
            } else {
                v3 = (v3 + "a")
                v10 := []int{fi7}
                _ = v10
            }
        } else {
            _ = len(v1)
            v3 = v3
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            switch (((((("\n\t" + (((v3 + "\n\t") + (v3 + "foo bar")) + v3)) + "foo bar") + v3) + "foo bar") + v3) + ((v3 + ((v3 + v3) + "\n\t")) + ("foo bar" + v3))) {
            case "":
                var v12 int16 = 42
                _ = v12
            case "a":
                v3 = (v3 + "foo bar")
            case "foo bar":
                var v13 uint = 255
                _ = v13
            }
            var v14 int64 = 1
            v15 := GenS1{F0: 2, F1: 114}
            var v16 int16 = -814
            if true {
                fmt.Println(v16, fi11, v16)
            } else {
                switch -2147483648 {
                case 329:
                    v17 := GenS1{F0: 42, F1: 18446744073709551615}
                    _ = v17
                case 128:
                    v3 = ((v3 + v3) + v3)
                }
            }
            v18 := []int{32767, 1, fi7, fi7}
            _ = fi11
            _ = v14
            _ = v15
            _ = v16
            _ = v18
        }
        _ = fi7
    }
    v19 := GenS1{F0: 4294967295, F1: 18446744073709551615}
    var v20 float32 = float32(0.0)
    v21 := GenS1{F0: 18446744073709551615, F1: 1}
    if true {
        var v22 int32 = -1
        v20 -= float32(0.1)
        v23 := []int{}
        _ = v22
        _ = v23
    }
    v3 = "hello"
    switch v3 {
    case "foo bar":
        var v24 int64 = 9223372036854775807
        _ = v24
    default:
        var v25 string = v3
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 bool = false
        var v28 uint64 = 1
        delete(v2, 2)
        _ = fi26
        _ = v27
        _ = v28
    }
    v3 = "foo bar"
    fmt.Println(v20, v20)
    var v29 int16 = -32768
    delete(v1, -1)
    v30 := GenS1{F0: 18446744073709551615, F1: 0}
    switch "\n\t" {
    case "":
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 float64 = 1e10
            v29 = -32768
            v29 = (v29 + (v29 % ((v29) | 1)))
            _ = fi31
            _ = v32
        }
    default:
        fmt.Println(v20, v29)
    }
    switch ("a" + (v3 + (v3 + ("hello" + v3)))) {
    case "a":
        v20 += (v20 / (float32(1e10) + v20))
    case "":
        var v33 uint = 2
        _ = v33
    case "foo bar":
        switch 1 {
        case 1:
            v29 = v29
        case -172:
            switch (v3 + v3) {
            case "":
                for fi34 := 0; fi34 < 10; fi34++ {
                    var v35 uint = 731
                    if (-1 == -633) {
                        var v36 uint32 = 0
                        v37 := []int{2147483647}
                        var v38 uint = uint(v36)
                        _ = v36
                        _ = v37
                        _ = v38
                    } else {
                        fmt.Println(fi34, fi34)
                        v35 = v35
                        v39 := GenS1{F0: 4294967295, F1: (^(v35))}
                        _ = v39
                    }
                    switch int(v29) {
                    case -9223372036854775808:
                        v40 := GenS2{F0: 642, F1: 4294967295}
                        _ = v40
                    case 32767:
                        v1[2147483647] = fi34
                    default:
                        fmt.Println(v3, v35)
                    }
                    _ = fi34
                    _ = v35
                }
            }
        }
    default:
        v41 := []int{}
        _ = v41
    }
    v42 := GenS3{F0: 0, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v42
}
