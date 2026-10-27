package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0() int16 {
    var v7 uint = 703
    fmt.Println(v7)
    v7 = v7
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int16 = -8
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 float64 = 0.1
    var v3 float32 = 1e10
    var v4 bool = false
    var v5 string = "a"
    var v6 int16 = 42
    v5 = (v5 + v5)
    if true {
        v11 := GenS1{F0: 1}
        switch 1 {
        case 4294967295:
            v12 := []int{}
            _ = v12
        case 2147483647:
            if v4 {
                fmt.Println(v5, v1)
                v1 -= (v1 >> ((956) & 63))
            }
        case 19:
            if v4 {
                v3 += (float32((float64((v3 + 1e10)) + float64(float32(float64(v3))))) * (v3 * v3))
                v13 := []int{-2147483648, -503, 32767}
                _ = v13
            }
        }
        _ = v11
    }
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := []int{(-(fi15)), fi15, -111, int(v1)}
        for fi17 := 0; fi17 < 5; fi17++ {
            v3 = v3
            v3 -= float32(1.5)
            var v18 int8 = -128
            if (true && (!((^(v14)) > 2))) {
                var v19 uint = (^(v1))
                if false {
                    if v4 {
                        var v20 string = ("\n\t" + (("foo bar" + "") + (v5 + "hello")))
                        _ = v20
                    }
                    if ((v18 + v18) >= v18) {
                        var v21 uint8 = 42
                        _ = v21
                    } else {
                        fmt.Println(v3, fi17, v18)
                        var v22 float32 = v3
                        _ = v22
                    }
                }
                var v23 float32 = (float32((v2 / (v2 - ((((float64(v3) * (float64(0.0) * (0.0 + (((v2 * (v2 * v2)) + (float64(v3) + 0.1)) * (v2 + v2))))) + v2) * (v2 * v2)) + float64(float32(0.1)))))) * v3)
                _ = v19
                _ = v23
            }
            _ = fi17
            _ = v18
        }
        v24 := GenS1{F0: 335}
        _ = fi15
        _ = v16
        _ = v24
    }
    if (100 >= 127) {
        if v4 {
            v6 = 1
        }
        fmt.Println(v14, v5, v3)
    }
    var v25 uint8 = 128
    var v26 uint64 = 4294967295
    fmt.Println(v2, v1)
    switch ("foo bar" + (v5 + "foo bar")) {
    case "hello":
        var v27 int16 = (v6 % (((v6 % ((v6) | 1))) | 1))
        _ = v27
    default:
        fmt.Println(v5, v5)
    }
    fmt.Println(v4)
    if ((v6 | (^(v6))) == (v6 * v6)) {
        if ((!v4) && v4) {
            fmt.Println(v26, v2, v4)
        } else {
            for fi28 := 0; fi28 < 15; fi28++ {
                for fi29 := 0; fi29 < 11; fi29++ {
                    if v4 {
                        var v30 int = fi28
                        v14 += (v14 - v14)
                        v5 = v5
                        v14 -= (^(v14))
                        _ = v30
                    } else {
                        v31 := GenS1{F0: -494}
                        v4 = (!((0 <= -438) || (!(false || false))))
                        _ = v31
                    }
                    v2 = (3.14159 / 0.1)
                    var v32 float32 = ((float32(3.14159) - 0.1) * (v3 * float32((1e10 + v2))))
                    _ = fi29
                    _ = v32
                }
                fmt.Println(fi28)
                _ = fi28
            }
        }
    }
    var v33 int16 = (v6 * v6)
    var v34 float64 = (1e10 - (3.14159 - float64((v3 + v3))))
    var v35 uint64 = 234
    switch v5 {
    case "foo bar":
        fmt.Println(v6, v14)
    case "\n\t":
        v1 = v1
    case "hello":
        v34 += float64(((v3 / v3) / v3))
    }
    switch (v5 + ("\n\t" + v5)) {
    case "a":
        if false {
            v14 = (v14 % ((32767) | 1))
        } else {
            v36 := GenS2{F0: (("foo bar" + v5) + v5)}
            _ = v36
        }
    case "hello":
        fmt.Println(v34, v2)
    default:
        for fi37 := 0; fi37 < 1; fi37++ {
            if (-9223372036854775808 != 9223372036854775807) {
                for fi38 := 0; fi38 < 10; fi38++ {
                    switch v5 {
                    case "\n\t":
                        v4 = ((v25 - (^(v25))) > 128)
                    }
                    v39 := []int{(fi38 * (fi37 - 32767)), (fi37 * (-(fi38))), (fi38 >> ((v1) & 63)), (fi38 << ((810) & 63))}
                    _ = fi38
                    _ = v39
                }
                var v40 uint16 = 42
                _ = v40
            }
            var v41 int64 = 1
            var v42 uint32 = 841
            _ = fi37
            _ = v41
            _ = v42
        }
    }
    v4 = (!v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
