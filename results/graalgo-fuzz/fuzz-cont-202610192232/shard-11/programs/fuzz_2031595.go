package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 int16) int {
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 += 70
        switch ("\n\t" + ("foo bar" + ("" + "hello"))) {
        case "":
            v6 := GenS1{F0: ("hello" + "")}
            _ = v6
        }
        var v7 uint16 = 65535
        _ = fi5
        _ = v7
    }
    p0 -= p0
    v8 := []int{866, 32767, -494}
    _ = p0
    _ = p1
    _ = v8
    return -62
}
func genFn1() int16 {
    v9 := []int{1, -9223372036854775808, genFn0(791, 749)}
    var v10 uint8 = 142
    v10 += 202
    v10 = v10
    _ = v10
    _ = v9
    return -920
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]int)
    var v3 uint8 = 42
    v4 := []int{}
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS2{F0: v3}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 uint16 = 42
    for fi15 := 0; fi15 < 7; fi15++ {
        switch ("foo bar" + v1) {
        case "":
            if (true || false) {
                if (387 != 240) {
                    var v16 int16 = -646
                    v3 -= 1
                    switch "\n\t" {
                    case "hello":
                        var v17 uint32 = 1
                        _ = v17
                    case "":
                        v3 -= (^(v3))
                    default:
                        var v18 uint8 = v3
                        _ = v18
                    }
                    _ = v16
                } else {
                    if false {
                        v19 := GenS2{F0: 255}
                        v1 = v1
                        for fi20 := 0; fi20 < 2; fi20++ {
                            var v21 int = -146
                            var v22 int32 = -2147483648
                            v23 := GenS3{F0: int(v14)}
                            fmt.Println(v21, fi20, v21)
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = v19
                    }
                    for fi24 := 0; fi24 < 1; fi24++ {
                        var v25 uint16 = 830
                        _ = fi24
                        _ = v25
                    }
                }
                var v26 uint16 = v14
                _ = v26
            } else {
                v14 = (v14 ^ (v14 / (((v14 >> ((4294967295) & 63))) | 1)))
                _ = len(v4)
                var v27 int = (fi15 + 32767)
                _ = v27
            }
        }
        _ = fi15
    }
    v1 = ("" + v1)
    v14 = 110
    switch 877 {
    case 2147483647:
        switch -226 {
        case -2147483648:
            fmt.Println(v3, v3, v1)
        case 0:
            switch v1 {
            case "":
                fmt.Println(v1)
            case "hello":
                v1 = v1
            default:
                var v28 uint64 = 2
                _ = v28
            }
        }
    case 42:
        for fi29 := 0; fi29 < 13; fi29++ {
            fmt.Println(v3, v1, v14)
            _ = fi29
        }
    case 9223372036854775807:
        v2[2147483648] = genFn0(1, -1)
    default:
        var v30 uint64 = 2
        _ = v30
    }
    if false {
        var v31 bool = (true || (!(false || (!(627 < genFn1())))))
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 string = ("foo bar" + v1)
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        v14 = 32767
        _ = fi34
    }
    var v35 string = ("a" + v1)
    switch (v1 + ((v35 + v35) + v35)) {
    case "a":
        v1 = "a"
    case "\n\t":
        var v36 uint64 = 562
        _ = v36
    case "":
        var v37 uint8 = v3
        _ = v37
    }
    v38 := GenS3{F0: -1}
    fmt.Println(v3, v35)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    delete(v2, 130)
    var v39 int64 = 42
    v14 = v14
    v14 += v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
}
