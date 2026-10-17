package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
func genFn0(p0 int) uint64 {
    v7 := []int{}
    switch "a" {
    case "":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = p0
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int32, p1 int) string {
    if ("a" >= "a") {
        v8 := []int{p1, p1, p1}
        _ = v8
    } else {
        if false {
            v9 := []int{(p1 - 2147483647), (p1 % ((32767) | 1))}
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int16 = 42
                fmt.Println(fi10, fi10)
                v11 -= (v11 ^ v11)
                _ = fi10
                _ = v11
            }
            if true {
                var v12 uint = 1
                var v13 uint = 4294967295
                if (1e10 == (1.5 + 0.1)) {
                    var v14 int32 = -1
                    var v15 uint16 = 65535
                    if ((v12 + uint(v15)) != (v12 + (^(v13)))) {
                        var v16 uint = (v12 % ((235) | 1))
                        _ = v16
                    }
                    _ = v14
                    _ = v15
                }
                _ = v12
                _ = v13
            }
            _ = v9
        } else {
            var v17 int16 = -375
            switch "" {
            case "":
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := GenS1{F0: (v17 / ((42) | 1)), F1: 32767, F2: (("" + ("foo bar" + ("hello" + ("\n\t" + ("\n\t" + ("foo bar" + "a")))))) + ("a" + (("" + "a") + (("foo bar" + ("\n\t" + "hello")) + ("hello" + ("a" + "foo bar"))))))}
                    fmt.Println(v17, v17)
                    _ = fi18
                    _ = v19
                }
            case "a":
                var v20 uint32 = 65535
                _ = v20
            default:
                p0 = p0
            }
            v21 := []int{4294967295, (p1 ^ 42), -611}
            _ = v17
            _ = v21
        }
        fmt.Println(p1, p0)
    }
    if (!true) {
        var v22 uint8 = 127
        var v23 uint32 = 65535
        if (0 >= uint16(v23)) {
            v24 := GenS1{F0: 325, F1: 42, F2: ""}
            _ = v24
        }
        _ = v22
        _ = v23
    } else {
        if false {
            p0 += (p0 % ((0) | 1))
            var v25 uint32 = 2147483648
            _ = v25
        }
    }
    var v26 uint64 = 0
    _ = p0
    _ = p1
    _ = v26
    return (("foo bar" + ("foo bar" + "\n\t")) + "foo bar")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 127
    var v3 int = (v2 / ((9223372036854775807) | 1))
    var v4 uint32 = 2147483647
    var v5 uint64 = 42
    var v6 float64 = (1e10 + 1e10)
    fmt.Println(v3, v4)
    var v27 float32 = ((1.5 / 1.5) * 0.1)
    v28 := []int{}
    v5 = v5
    delete(v1, v4)
    var v29 uint = 612
    v4 = v4
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    v27 -= float32(float64(3.14159))
    fmt.Println(v27)
    switch (-(v2)) {
    case -9223372036854775808:
        switch v3 {
        case -2147483648:
            var v30 int8 = 0
            _ = v30
        default:
            v5 = genFn0(v3)
        }
    default:
        if (v6 >= (v6 - v6)) {
            v31 := GenS1{F0: 1, F1: 615, F2: ("" + (("hello" + ("foo bar" + "hello")) + "a"))}
            switch ("" + "foo bar") {
            case "a":
                v6 = v6
            }
            v32 := GenS1{F0: 0, F1: 0, F2: "hello"}
            _ = v31
            _ = v32
        }
    }
    var v33 float64 = float64(3.14159)
    var v34 int16 = 731
    v3 = int(v29)
    v6 -= 1e10
    fmt.Println(v34)
    v33 += v33
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
