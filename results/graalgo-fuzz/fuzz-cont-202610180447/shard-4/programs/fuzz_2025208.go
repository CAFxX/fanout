package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
func genFn0(p0 uint8, p1 uint, p2 string) int32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int8 = 89
        v6 := []int{42, -9223372036854775808, fi4}
        p0 += 24
        _ = fi4
        _ = v5
        _ = v6
    }
    p1 -= 65535
    var v7 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}
func genFn1() string {
    var v8 int8 = 100
    if (1.5 >= 1e10) {
        v9 := []int{9223372036854775807, 127, 216}
        fmt.Println(v8)
        _ = v9
    }
    var v10 bool = (2147483647 >= 2147483647)
    var v11 int32 = 42
    _ = v10
    _ = v11
    _ = v8
    return "foo bar"
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    for fi12 := 0; fi12 < 19; fi12++ {
        v3 -= v3
        _ = fi12
    }
    v13 := GenS1{F0: -19, F1: ("\n\t" + "foo bar")}
    fmt.Println(v1, v1)
    var v14 string = ("hello" + ("foo bar" + ""))
    v15 := GenS1{F0: 18, F1: v14}
    fmt.Println(v1, v14)
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            fmt.Println(v14, v3)
        } else {
            var v17 bool = (!false)
            _ = len(v2)
            var v18 int32 = genFn0(v3, 1, ("a" + v14))
            _ = v17
            _ = v18
        }
        if (0 < 2147483648) {
            fmt.Println(v14)
            fmt.Println(v14, v14)
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int8 = -83
                v1 = genFn0(1, 42, v14)
                switch ("hello" + "hello") {
                case "":
                    v2[(v14 + (("hello" + "\n\t") + ("foo bar" + "a")))] = 2
                    switch ("" + v14) {
                    case "":
                        switch fi16 {
                        case 4294967295:
                            v3 = (v3 + v3)
                        case 128:
                            v21 := []int{}
                            v21 = append(v21, -1)
                            fmt.Println(fi19, v14, fi19)
                            _ = v21
                        case 9223372036854775807:
                            v22 := []int{fi16, fi16, fi16, (fi16 & fi19)}
                            v14 = (("\n\t" + "foo bar") + "foo bar")
                            v14 = v14
                            _ = v22
                        default:
                            var v23 uint32 = 0
                            _ = v23
                        }
                    case "foo bar":
                        var v24 int8 = v20
                        _ = v24
                    }
                default:
                    v25 := []int{}
                    _ = v25
                }
                _ = fi19
                _ = v20
            }
        } else {
            v14 = v14
            var v26 int16 = -32768
            var v27 int8 = 17
            _ = v26
            _ = v27
        }
        _ = fi16
    }
    switch (((((v14 + "\n\t") + (v14 + (v14 + v14))) + "") + v14) + ("a" + (("foo bar" + "a") + v14))) {
    case "hello":
        v28 := GenS1{F0: 0, F1: v14}
        _ = v28
    case "a":
        fmt.Println(v3, v14, v14)
    case "\n\t":
        delete(v2, "a")
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        fmt.Println(v1, fi29, v14)
        _ = fi29
    }
    if false {
        v1 -= -1
        v14 = v14
    } else {
        v30 := GenS2{F0: 516, F1: (v1 * (v1 ^ (v1 * -2147483648))), F2: 9223372036854775807}
        var v31 uint32 = 42
        for fi32 := 0; fi32 < 1; fi32++ {
            fmt.Println(v1, fi32, v3)
            var v33 int16 = -1
            var v34 int32 = v1
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v30
        _ = v31
    }
    if (false && true) {
        v35 := []int{1}
        _ = v35
    }
    v36 := GenS2{F0: 0, F1: genFn0((^(v3)), 18446744073709551615, "hello"), F2: 563}
    for fi37 := 0; fi37 < 10; fi37++ {
        for fi38 := 0; fi38 < 3; fi38++ {
            fmt.Println(fi38, fi37, v1)
            v39 := GenS3{F0: 127, F1: (fi37 << ((18446744073709551615) & 63)), F2: 954}
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
}
