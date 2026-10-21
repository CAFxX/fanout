package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 int32, p2 int32) string {
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 -= p0
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    p0 = 100
    var v10 string = ("\n\t" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return ((((v10 + v10) + v10) + (v10 + v10)) + v10)
}
func genFn1() int32 {
    var v11 int8 = 100
    v11 += (v11 + (v11 << ((266) & 63)))
    v11 -= -1
    _ = v11
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint64 = 668
    v4 := make(map[string]string)
    v5 := []string{("\n\t" + (("\n\t" + "foo bar") + ("foo bar" + ""))), "hello", "foo bar"}
    var v6 bool = (-1 < -815)
    var v12 string = "hello"
    if (0 == -80) {
        delete(v1, -2147483648)
        fmt.Println(v6)
        var v13 int64 = -1
        _ = v13
    } else {
        if (0.1 > float32((float64(-2.25) - 1.5))) {
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 float32 = 0.1
                _ = fi14
                _ = v15
            }
            var v16 uint32 = 663
            switch v12 {
            case "":
                for fi17 := 0; fi17 < 10; fi17++ {
                    v18 := GenS1{F0: v16}
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 uint16 = 42
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        _ = fi19
                        _ = v20
                    }
                    v21 := v4[v12]
                    _ = fi17
                    _ = v18
                    _ = v21
                }
            case "a":
                if (!(127 > 100)) {
                    v22 := v1[-1]
                    _ = v22
                } else {
                    v2[v16] = v12
                    var v23 int32 = -2147483648
                    _ = v23
                }
            }
            _ = v16
        } else {
            var v24 int8 = 0
            v1[821] = (247 < -1)
            v12 = ("\n\t" + v12)
            _ = v24
        }
        if v6 {
            var v25 uint64 = v3
            _ = v25
        }
        var v26 int32 = 51
        _ = v26
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        if (32767 >= 1) {
            v3 += (^(v3))
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 int16 = 0
            switch "\n\t" {
            case "foo bar":
                v29 = 0
            case "hello":
                v30 := v2[194]
                _ = v30
            case "":
                fmt.Println(v3, v3)
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v31 uint64 = 0
    v6 = (-2.25 > 1e10)
    fmt.Println(v6)
    var v32 bool = true
    v6 = (1000 <= -867)
    if ((42 > -128) || v6) {
        fmt.Println(v32)
        v33 := GenS2{F0: "\n\t"}
        _ = v33
    }
    var v34 int8 = -82
    v34 = v34
    v3 = (v3 >> ((65535) & 63))
    v35 := GenS2{F0: "a"}
    if (32767 == 0) {
        v36 := GenS2{F0: "a"}
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v38 := v1[fi37]
        v39 := v1[fi37]
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 bool = (255 > 209)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
