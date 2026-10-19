package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int16; F2 int64 }
func genFn0(p0 uint16, p1 uint16, p2 int32) string {
    var v6 int32 = p2
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            v6 -= 42
            var v8 uint16 = p1
            for fi9 := 0; fi9 < 6; fi9++ {
                fmt.Println(p2, p1, fi7)
                switch (fi7 / (((fi7 - fi9)) | 1)) {
                case 32767:
                    if ("hello" > "a") {
                        v10 := []int{fi9, 1}
                        if (128 >= 0) {
                            if false {
                                if len(v10) > 0 {
                                    v10[0] = fi7
                                }
                                v8 -= (v8 & (p1 << ((42) & 63)))
                            }
                            fmt.Println(v6)
                            var v11 int16 = -32768
                            _ = v11
                        }
                        if false {
                            v12 := GenS1{F0: ((("foo bar" + "foo bar") + "\n\t") + "")}
                            if len(v10) > 0 {
                                v10[0] = (-(fi9))
                            }
                            _ = v12
                        } else {
                            switch "\n\t" {
                            case "a":
                                p0 -= p1
                            case "hello":
                                fmt.Println(p1)
                            case "":
                                if len(v10) > 0 {
                                    v10[0] = fi7
                                }
                            }
                        }
                        _ = v10
                    } else {
                        v13 := []int{fi7, fi7}
                        var v14 string = (("\n\t" + ((("hello" + "\n\t") + "\n\t") + "")) + "foo bar")
                        var v15 uint32 = 2147483648
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                case 32768:
                    v16 := GenS1{F0: "hello"}
                    _ = v16
                default:
                    v17 := []int{fi7, fi9}
                    _ = v17
                }
                v18 := []int{}
                _ = fi9
                _ = v18
            }
            _ = fi7
            _ = v8
        }
    } else {
        if false {
            var v19 string = "a"
            _ = v19
        } else {
            if true {
                for fi20 := 0; fi20 < 10; fi20++ {
                    v6 = (v6 ^ v6)
                    _ = fi20
                }
            }
        }
        p1 = 756
        var v21 uint32 = 2147483648
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "a"
}

func main() {
    var v1 bool = (965 <= 925)
    v2 := []int{4294967295}
    v3 := []uint32{1, 503, 381, 467}
    var v4 uint32 = 2
    var v5 float32 = 1.5
    if false {
        for fi22 := 0; fi22 < 10; fi22++ {
            v5 += float32(0.1)
            _ = fi22
        }
    }
    v5 += float32(3.14159)
    v3 = append(v3, (v4 % ((v4) | 1)))
    if v1 {
        var v23 int64 = -801
        v23 -= v23
        _ = v23
    }
    if (28 >= -64) {
        v5 -= 0.0
        for fi24 := 0; fi24 < 13; fi24++ {
            fmt.Println(v1, fi24)
            v25 := GenS2{F0: 526, F1: 32767, F2: 568}
            _ = fi24
            _ = v25
        }
        switch (((("\n\t" + "foo bar") + ("a" + "")) + "a") + ("a" + "foo bar")) {
        case "hello":
            v4 = (v4 * v4)
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = 127
            }
        case "a":
            v1 = ((v1 && ((v4 * (v4 ^ (v4 ^ 937))) <= v4)) || (false || true))
        }
    } else {
        var v26 int32 = -2147483648
        _ = v26
    }
    v27 := GenS2{F0: 0, F1: -502, F2: 0}
    for fi28 := 0; fi28 < 14; fi28++ {
        if len(v2) > 0 {
            v2[0] = (fi28 * fi28)
        }
        _ = fi28
    }
    var v29 float32 = v5
    var v30 uint = 785
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
