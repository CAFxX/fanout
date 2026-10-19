package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    if (42 > 2147483648) {
        switch -1 {
        case -1:
            if (42 == 647) {
                var v6 uint64 = 1
                v7 := GenS1{F0: ("foo bar" + "\n\t"), F1: 232}
                v6 = (^(v6))
                _ = v6
                _ = v7
            } else {
                switch ("a" + "\n\t") {
                case "hello":
                    v8 := GenS2{F0: "foo bar"}
                    _ = v8
                }
            }
        default:
            v9 := GenS2{F0: ""}
            _ = v9
        }
        var v10 uint = 205
        _ = v10
    }
    var v11 uint64 = 873
    v12 := []int{9223372036854775807}
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 2
    v3 := make(map[string]int)
    var v4 uint32 = 4294967295
    var v5 float64 = 1e10
    var v13 int64 = -9223372036854775808
    if true {
        fmt.Println(v4, v2, v5)
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            fmt.Println(v13, fi14)
            var v15 int = fi14
            for fi16 := 0; fi16 < 18; fi16++ {
                if false {
                    _ = len(v3)
                    if true {
                        v4 -= (v4 | 4294967295)
                    }
                } else {
                    var v17 bool = (v5 == v5)
                    _ = v17
                }
                var v18 int16 = -1
                _ = fi16
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        delete(v1, "\n\t")
    }
    var v19 bool = false
    if ((4294967295 <= 1) || true) {
        var v20 int16 = 1000
        v21 := GenS2{F0: ("" + "")}
        v13 = genFn0()
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS1{F0: "\n\t", F1: (^(v13))}
                delete(v1, "foo bar")
                v25 := v3["hello"]
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
    }
    v26 := GenS3{F0: 65535}
    fmt.Println(v4, v2)
    fmt.Println(v4, v4, v13)
    var v27 int16 = 32767
    switch ((("hello" + "hello") + "") + ("hello" + "")) {
    case "foo bar":
        v28 := GenS3{F0: (v4 % ((v4) | 1))}
        _ = v28
    case "":
        var v29 int32 = 2147483647
        _ = v29
    case "a":
        _ = len(v3)
    }
    _ = len(v3)
    if v19 {
        for fi30 := 0; fi30 < 2; fi30++ {
            v31 := []int{(fi30 << ((308) & 63)), (-(fi30)), (fi30 & fi30), fi30}
            _ = fi30
            _ = v31
        }
        v2 += 2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
