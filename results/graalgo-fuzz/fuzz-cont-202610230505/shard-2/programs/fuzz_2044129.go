package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint32 }
func genFn0() int16 {
    var v5 int8 = 42
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(fi6, v5)
            if false {
                var v8 uint8 = 229
                var v9 string = "hello"
                v8 = (v8 + 1)
                _ = v8
                _ = v9
            } else {
                var v10 uint64 = 4294967295
                v10 -= 0
                _ = v10
            }
            _ = fi7
        }
        _ = fi6
    }
    _ = v5
    return 42
}
func genFn1(p0 uint16, p1 uint8, p2 uint8) bool {
    for fi11 := 0; fi11 < 1; fi11++ {
        if (true || (("a" + ("hello" + ("hello" + ("\n\t" + "a")))) == "\n\t")) {
            if true {
                var v12 float32 = (1.5 - ((float32((-2.25 / 3.14159)) / 0.1) + float32((float64(1.5) - 1e10))))
                var v13 float32 = v12
                _ = v12
                _ = v13
            } else {
                var v14 int64 = -9223372036854775808
                fmt.Println(p2, fi11)
                _ = v14
            }
            v15 := []int{(fi11 - fi11), fi11}
            switch "" {
            case "a":
                var v16 uint64 = 153
                _ = v16
            case "foo bar":
                p2 = 152
            case "":
                for fi17 := 0; fi17 < 2; fi17++ {
                    if len(v15) > 0 {
                        v15[0] = 128
                    }
                    fmt.Println(p1, p1, fi17)
                    _ = fi17
                }
            default:
                v18 := GenS1{F0: -2147483648, F1: 968, F2: 1000}
                _ = v18
            }
            _ = v15
        }
        _ = fi11
    }
    var v19 int8 = 0
    v20 := GenS1{F0: 311, F1: 0, F2: 90}
    if (-2147483648 != -1) {
        if false {
            v19 = v19
            var v21 string = "foo bar"
            v22 := []int{-548, -173, 9223372036854775807, 566}
            _ = v21
            _ = v22
        } else {
            switch -2147483648 {
            case 128:
                switch 0 {
                case -289:
                    var v23 uint = 1
                    _ = v23
                case 42:
                    p1 += (p1 & (p1 ^ (p1 % ((p2) | 1))))
                default:
                    v24 := GenS1{F0: 223, F1: 2147483648, F2: -425}
                    _ = v24
                }
            }
        }
        p2 += 2
        if true {
            var v25 int8 = v19
            _ = v25
        } else {
            p2 = p2
            if (!(int16(p2) == 1)) {
                p2 -= p1
                v26 := GenS1{F0: 42, F1: 2, F2: 0}
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return true
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int64 = 1
    var v4 string = ("foo bar" + ("" + ""))
    for fi27 := 0; fi27 < 7; fi27++ {
        v4 = ((v4 + "") + "hello")
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 int32 = 42
            v30 := GenS1{F0: (v29 ^ v29), F1: 436, F2: 42}
            var v31 float64 = (1.5 * (-2.25 * 1e10))
            v2[2147483647] = true
            switch (fi27 ^ (fi27 - fi27)) {
            case 32767:
                v31 += ((3.14159 - (v31 - v31)) - ((1e10 * (3.14159 / 1.0)) * (3.14159 * (1e10 + (float64(0.0) * (((v31 * float64(1e10)) / (float64((float32(v31) + (-2.25 / 1.5))) - v31)) - v31))))))
            }
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
        _ = fi27
    }
    switch 666 {
    case 32768:
        fmt.Println(v3)
    }
    if false {
        fmt.Println(v3, v3)
        v32 := []int{}
        _ = v32
    }
    if false {
        var v33 uint16 = 287
        fmt.Println(v33)
        v33 = v33
        _ = v33
    }
    if true {
        v3 += v3
    }
    v34 := GenS2{F0: -1}
    if (false && (true || (!false))) {
        v35 := GenS3{F0: 2, F1: -767, F2: 721}
        _ = v35
    }
    v36 := []int{2147483647, 0}
    v4 = "hello"
    var v37 string = (v4 + "foo bar")
    var v38 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
