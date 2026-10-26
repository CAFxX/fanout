package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 int32) int16 {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = (-9223372036854775808 != 1)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -9223372036854775808}
    for fi10 := 0; fi10 < 8; fi10++ {
        p0 = 42
        _ = fi10
    }
    _ = p0
    _ = v9
    return 32767
}
func genFn1(p0 uint32, p1 bool) uint16 {
    if p1 {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := []int{313, (fi12 ^ (^(fi12))), fi11, -1}
                var v14 int64 = -418
                _ = fi12
                _ = v13
                _ = v14
            }
            if p1 {
                if (false || true) {
                    p0 -= p0
                } else {
                    p1 = true
                    fmt.Println(p1, p0)
                    fmt.Println(fi11, fi11)
                }
                var v15 int32 = 574
                if (false && (-4 > -112)) {
                    var v16 uint16 = 534
                    var v17 uint32 = 0
                    _ = v16
                    _ = v17
                } else {
                    if (!p1) {
                        var v18 int32 = (v15 >> ((179) & 63))
                        v18 = 2147483647
                        _ = v18
                    }
                    switch fi11 {
                    case 311:
                        var v19 uint64 = 42
                        _ = v19
                    case 32767:
                        var v20 uint16 = 2
                        _ = v20
                    }
                    p0 += p0
                }
                _ = v15
            }
            if false {
                if (1e10 < 3.14159) {
                    p0 += (p0 * 0)
                    var v21 uint = 0
                    _ = v21
                }
            } else {
                p1 = (!true)
            }
            _ = fi11
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := []int{fi22, fi22}
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int8 = 47
                _ = fi24
                _ = v25
            }
            switch fi22 {
            case 128:
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 uint32 = 2147483648
                    _ = fi26
                    _ = v27
                }
            }
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 int = 128
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 int8 = 0
    v5 := []string{"hello", (("foo bar" + ("" + "a")) + ("foo bar" + "foo bar"))}
    v6 := make(map[string]int)
    var v28 uint32 = uint32(v3)
    if ((42 > 42) || ("\n\t" <= "a")) {
        var v29 uint16 = 255
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        if (-809 == -885) {
            var v31 string = "hello"
            var v32 uint16 = 255
            v3 += v3
            _ = v31
            _ = v32
        }
        _ = fi30
    }
    v3 -= v3
    if true {
        v33 := GenS2{F0: 1, F1: 7}
        v34 := GenS2{F0: -1, F1: 44}
        var v35 uint8 = 0
        _ = v33
        _ = v34
        _ = v35
    }
    if (true && ((0.1 / 3.14159) != -2.25)) {
        for fi36 := 0; fi36 < 20; fi36++ {
            fmt.Println(v4, v1)
            var v37 int = (fi36 + 478)
            _ = fi36
            _ = v37
        }
        fmt.Println(v1)
        if len(v5) > 0 {
            v5[0] = (("hello" + ("" + (("\n\t" + ("\n\t" + "foo bar")) + "\n\t"))) + ("a" + ""))
        }
    } else {
        fmt.Println(v3)
    }
    var v38 float32 = 1.5
    v39 := v2[(v1 ^ 9223372036854775807)]
    for fi40 := 0; fi40 < 1; fi40++ {
        if len(v5) > 0 {
            v5[0] = ("hello" + "")
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
