package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0() string {
    var v6 string = ""
    v6 = ""
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    v9 := []int{32768, 9223372036854775807, 152, 1}
    _ = v6
    _ = v9
    return ("\n\t" + (v6 + (v6 + v6)))
}
func genFn1(p0 uint32, p1 uint64) int {
    p0 += (p0 % ((p0) | 1))
    p1 = (p1 & (p1 << ((0) & 63)))
    for fi10 := 0; fi10 < 10; fi10++ {
        if (1e10 == (float64(0.1) - (float64(3.14159) + 3.14159))) {
            switch (^(fi10)) {
            case -1:
                fmt.Println(p0)
            case -243:
                for fi11 := 0; fi11 < 10; fi11++ {
                    if (p1 == p1) {
                        switch (fi10 - 128) {
                        case 98:
                            if false {
                                switch ("a" + (("a" + (("hello" + ("a" + ("hello" + "hello"))) + "\n\t")) + ("foo bar" + "a"))) {
                                case "":
                                    v12 := []int{fi11, (-(fi11))}
                                    _ = v12
                                case "\n\t":
                                    p0 += p0
                                case "hello":
                                    v13 := []int{-9223372036854775808, (fi11 | 32768)}
                                    _ = v13
                                default:
                                    v14 := GenS1{F0: 42, F1: 255, F2: (7 >= int8(p0))}
                                    _ = v14
                                }
                            } else {
                                p1 -= (p1 * 18446744073709551615)
                                p1 += p1
                                var v15 int32 = 0
                                _ = v15
                            }
                        }
                        fmt.Println(fi11, fi11, p1)
                    } else {
                        var v16 uint8 = 46
                        _ = v16
                    }
                    _ = fi11
                }
            case -596:
                var v17 int16 = 0
                _ = v17
            }
            p1 += p1
            var v18 bool = ((1000 <= 42) || false)
            _ = v18
        }
        if (uint(p0) >= 2) {
            var v19 bool = ((^(fi10)) <= (fi10 / (((fi10 / ((-224) | 1))) | 1)))
            switch fi10 {
            case 0:
                var v20 float32 = -2.25
                _ = v20
            case 2147483647:
                v21 := GenS1{F0: 399, F1: 916, F2: ((!true) || ((!(0 <= 42)) && (!("\n\t" >= "a"))))}
                _ = v21
            }
            _ = v19
        }
        fmt.Println(p0, p0)
        _ = fi10
    }
    _ = p0
    _ = p1
    return 723
}

func main() {
    var v1 bool = (42 > 9223372036854775807)
    v2 := make(map[string]uint32)
    var v3 int16 = 770
    var v4 float64 = float64((-2.25 / 1.5))
    var v5 int8 = 1
    var v22 int8 = -128
    delete(v2, ("foo bar" + ("foo bar" + "\n\t")))
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS2{F0: 0, F1: 42}
        _ = fi23
        _ = v24
    }
    var v25 int8 = (^(v22))
    var v26 uint32 = 212
    var v27 uint16 = 1
    if (1 != 0) {
        delete(v2, "foo bar")
        var v28 float32 = 1.5
        _ = v28
    }
    var v29 string = (("a" + ("hello" + (("foo bar" + "a") + "foo bar"))) + ("a" + ("a" + "\n\t")))
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := []int{(fi30 | (fi30 + (fi30 * genFn1(4294967295, 18446744073709551615)))), (fi30 - fi30), fi30}
        _ = fi30
        _ = v31
    }
    var v32 int64 = 0
    if (2147483647 == -1) {
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 int64 = v32
            var v35 int8 = v5
            _ = fi33
            _ = v34
            _ = v35
        }
        v36 := v2[v29]
        var v37 float32 = (1e10 * 1.5)
        _ = v36
        _ = v37
    }
    v38 := GenS3{F0: 2, F1: 0}
    if ((v32 << ((65535) & 63)) < 750) {
        v5 = 12
    }
    v39 := []int{-2147483648}
    switch v29 {
    case "\n\t":
        v32 += 42
    case "":
        v40 := GenS2{F0: v32, F1: 4294967295}
        _ = v40
    case "a":
        for fi41 := 0; fi41 < 1; fi41++ {
            v42 := GenS1{F0: (v27 / ((880) | 1)), F1: 4294967295, F2: v1}
            _ = fi41
            _ = v42
        }
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
