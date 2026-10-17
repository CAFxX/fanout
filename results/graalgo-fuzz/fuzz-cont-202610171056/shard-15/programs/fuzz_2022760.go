package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
func genFn0(p0 int8, p1 uint8, p2 string) int16 {
    if false {
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(fi7, fi7)
            var v8 float32 = 1.5
            _ = fi7
            _ = v8
        }
        var v9 int8 = 100
        if ((0.0 >= float64(3.14159)) || (true && true)) {
            if (1 < 2147483647) {
                v10 := []int{-86, -888, 9223372036854775807, 42}
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 bool = (-32768 <= 0)
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
        }
        _ = v9
    } else {
        v13 := GenS1{F0: 682, F1: false, F2: p1}
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: false, F2: (p1 ^ (p1 >> ((896) & 63)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 32767
}
func genFn1() int16 {
    var v15 uint64 = 4294967295
    fmt.Println(v15)
    var v16 uint32 = 0
    switch 32768 {
    case -1:
        var v17 float64 = float64(0.0)
        _ = v17
    case 42:
        fmt.Println(v15, v16)
    }
    _ = v15
    _ = v16
    return -495
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"foo bar"}
    v3 := []string{"a", ("\n\t" + (("a" + "hello") + "foo bar")), ((((("foo bar" + "a") + "a") + (("\n\t" + "foo bar") + (("a" + ("" + "foo bar")) + "foo bar"))) + "foo bar") + ("" + ((("\n\t" + "hello") + ("foo bar" + "\n\t")) + "hello"))), ""}
    v4 := []uint32{528, 1}
    var v5 string = "hello"
    var v6 uint32 = 401
    v5 = v5
    var v18 uint16 = 65535
    if true {
        v19 := v1[127]
        _ = v19
    }
    fmt.Println(v18, v6, v5)
    var v20 float32 = 3.14159
    v21 := v1[12]
    v20 += 1e10
    v21 = "\n\t"
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            v20 = v20
            v2 = append(v2, "a")
            if true {
                v21 = (((v21 + v5) + v21) + v21)
                v21 = v21
                v20 = v20
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v21 = (v21 + "")
        v25 := GenS1{F0: 528, F1: ((v18 & 958) != v18), F2: 255}
        _ = fi24
        _ = v25
    }
    fmt.Println(v21)
    if (!((!(2147483647 == 1)) && (int32(v18) == 1))) {
        if false {
            v20 = v20
            v20 = v20
        }
    } else {
        var v26 uint = 1
        v6 -= (^(v6))
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
