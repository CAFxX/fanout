package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 string; F1 string; F2 int }
func genFn0(p0 int8, p1 string, p2 uint32) int16 {
    var v4 string = p1
    var v5 int32 = -1
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 uint8, p2 int32) uint8 {
    fmt.Println(p0, p2)
    if false {
        var v6 int = -686
        v7 := GenS1{F0: 0, F1: (p2 | 42), F2: (((((0.0 - (0.1 - 1e10)) + (-2.25 - 1.5)) - float32(1.5)) != 1.5) || true)}
        p0 = (-(p0))
        _ = v6
        _ = v7
    }
    p1 += 0
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 float64 = (0.0 * 0.0)
    v2 := make(map[int]bool)
    var v3 float64 = v1
    if (2 > 42) {
        v2[-566] = (false && (false && (18446744073709551615 >= 0)))
    }
    if (false || (0 != 0)) {
        var v8 uint16 = 859
        _ = v8
    } else {
        if ((((3.14159 + (v1 / v3)) * (v3 - (v1 + v3))) > v3) && (true && false)) {
            if (3.14159 == -2.25) {
                v1 = v3
            }
        }
        if false {
            var v9 bool = false
            var v10 float64 = (v3 * float64(float32(v3)))
            var v11 float64 = float64(1.5)
            _ = v10
            _ = v11
            _ = v9
        } else {
            var v12 uint32 = 65535
            if false {
                v13 := GenS2{F0: "", F1: ("hello" + ((((((("" + "") + "foo bar") + ("a" + "")) + ("foo bar" + "")) + "foo bar") + "") + "a")), F2: 42}
                _ = v13
            } else {
                var v14 float64 = v3
                v15 := []int{32767, 1}
                _ = v14
                _ = v15
            }
            _ = v12
        }
        v16 := GenS1{F0: 0, F1: 984, F2: (42 <= 42)}
        _ = v16
    }
    v3 = v1
    var v17 float32 = float32(0.0)
    if (0 == 89) {
        var v18 int16 = 1
        v3 = v3
        for fi19 := 0; fi19 < 18; fi19++ {
            v2[-813] = ((false || true) && (false || (!(v18 < v18))))
            v18 = v18
            switch ("\n\t" + "") {
            case "a":
                v20 := GenS2{F0: "", F1: ("foo bar" + ((("hello" + (("hello" + "\n\t") + "\n\t")) + "hello") + "a")), F2: fi19}
                _ = v20
            default:
                fmt.Println(v18)
            }
            _ = fi19
        }
        _ = v18
    }
    fmt.Println(v1, v3, v3)
    for fi21 := 0; fi21 < 20; fi21++ {
        v3 -= float64(0.0)
        _ = fi21
    }
    v3 -= 1e10
    var v22 float32 = 0.0
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 uint64 = 2
            v1 += v3
            _ = fi24
            _ = v25
        }
        switch (fi23 - 778) {
        case -628:
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := []int{fi23, fi23, fi23, 32768}
                var v28 float64 = v1
                if ((false && true) || (-2147483648 < 2147483647)) {
                    v17 = v22
                } else {
                    v3 -= float64(v17)
                    v1 = ((1e10 - v3) * (v28 / (v3 * (v3 / v28))))
                    v29 := GenS1{F0: 18446744073709551615, F1: -899, F2: (true || ((!true) || (552 != 255)))}
                    _ = v29
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        case 568:
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 uint8 = 48
                v32 := GenS1{F0: 4294967295, F1: -764, F2: true}
                v3 = 1e10
                v33 := GenS2{F0: ("hello" + (("" + ("\n\t" + (("\n\t" + "foo bar") + "foo bar"))) + "\n\t")), F1: "foo bar", F2: (fi23 / ((42) | 1))}
                if true {
                    delete(v2, 127)
                    v31 += v31
                    fmt.Println(v1)
                }
                v17 = v17
                fmt.Println(v17, v17)
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        case -9223372036854775808:
            v22 = 3.14159
        default:
            v1 = ((1.5 - (1.5 * v1)) / (-2.25 + -2.25))
        }
        v3 -= float64((v22 * (float32((3.14159 - float64(v17))) + float32(v1))))
        _ = fi23
    }
    v22 = (3.14159 / v22)
    var v34 string = (("" + (("a" + ("" + ("foo bar" + (("a" + "") + "foo bar")))) + ("hello" + ("\n\t" + "")))) + ("\n\t" + "a"))
    var v35 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
}
