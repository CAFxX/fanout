package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int8 }
func genFn0() int16 {
    var v4 uint32 = 0
    switch 127 {
    case -465:
        var v5 float32 = (float32(0.1) * -2.25)
        _ = v5
    default:
        v4 += (^(v4))
    }
    if (100 >= 42) {
        var v6 string = "hello"
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := []int{}
            var v9 uint = 65535
            v6 = ("" + "")
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int64 = -868
        _ = v10
        _ = v6
    } else {
        v4 += 65535
        v11 := GenS1{F0: 128}
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(fi12, fi13)
            var v14 float32 = 3.14159
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if (128 == 2) {
        for fi15 := 0; fi15 < 18; fi15++ {
            switch fi15 {
            case 473:
                v4 -= 0
            case 970:
                var v16 string = ("" + (("foo bar" + "") + "a"))
                _ = v16
            default:
                v17 := GenS2{F0: 32767, F1: (v4 + v4)}
                _ = v17
            }
            v18 := []int{691, (fi15 - (fi15 | 25))}
            _ = len(v18)
            _ = fi15
            _ = v18
        }
    }
    _ = v4
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{1}
    var v3 int32 = 42
    for fi19 := 0; fi19 < 2; fi19++ {
        v3 = v3
        switch "a" {
        case "a":
            switch "foo bar" {
            case "":
                var v20 int32 = -406
                _ = v20
            case "foo bar":
                for fi21 := 0; fi21 < 6; fi21++ {
                    fmt.Println(fi19, fi19, fi19)
                    _ = fi21
                }
            case "\n\t":
                v3 += (-(v3))
            }
        }
        var v22 uint32 = 987
        _ = fi19
        _ = v22
    }
    v3 = -1
    v3 += v3
    v3 += (v3 >> ((18446744073709551615) & 63))
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            v3 = -2147483648
            _ = fi24
        }
        delete(v1, fi23)
        v25 := GenS1{F0: (fi23 + -1)}
        _ = fi23
        _ = v25
    }
    if true {
        if (!false) {
            v3 = v3
            fmt.Println(v3)
            v2 = append(v2, 9223372036854775807)
        }
        if false {
            fmt.Println(v3)
            var v26 int32 = (v3 ^ -430)
            if false {
                v1[-805] = 42
                v3 += v26
                if (false || false) {
                    var v27 uint8 = 127
                    var v28 int16 = -513
                    _ = v27
                    _ = v28
                } else {
                    switch ((("foo bar" + (("a" + ("\n\t" + ("hello" + "hello"))) + "\n\t")) + ((("foo bar" + ("hello" + ("hello" + "hello"))) + ("\n\t" + "a")) + ("\n\t" + "\n\t"))) + ("foo bar" + "\n\t")) {
                    case "foo bar":
                        v29 := GenS3{F0: 42, F1: -9223372036854775808, F2: 1}
                        _ = v29
                    case "\n\t":
                        fmt.Println(v26, v3)
                    }
                    fmt.Println(v26, v26)
                }
            } else {
                v3 = v26
                switch 2147483647 {
                case 701:
                    var v30 int64 = 9223372036854775807
                    _ = v30
                }
            }
            _ = v26
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            switch (^(fi31)) {
            case 4294967295:
                v3 = v3
            }
            _ = fi31
        }
    }
    if (0 >= 0) {
        v32 := GenS2{F0: 635, F1: 61}
        if (!true) {
            v33 := GenS4{F0: 157, F1: 65535}
            _ = v33
        }
        fmt.Println(v3)
        _ = v32
    } else {
        v3 -= (v3 % ((v3) | 1))
        v34 := GenS2{F0: -32768, F1: 42}
        delete(v1, 42)
        _ = v34
    }
    if (!(189 > 0)) {
        var v35 string = ""
        _ = v35
    } else {
        v3 += v3
        v36 := GenS5{F0: 42}
        _ = v36
    }
    var v37 int16 = -919
    var v38 uint = 84
    v1[int(v3)] = 32767
    var v39 float32 = -2.25
    fmt.Println(v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
