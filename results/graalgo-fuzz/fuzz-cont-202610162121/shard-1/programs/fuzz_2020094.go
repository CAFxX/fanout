package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int16, p1 int16, p2 int8) uint64 {
    p2 -= (p2 >> ((609) & 63))
    fmt.Println(p2, p2, p2)
    p1 = (p1 * (^(p1)))
    v4 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}
func genFn1(p0 bool, p1 int32) uint {
    for fi5 := 0; fi5 < 5; fi5++ {
        if (!true) {
            var v6 int64 = 1
            v6 = (v6 ^ v6)
            _ = v6
        } else {
            p1 -= 593
            v7 := GenS1{F0: 77}
            if true {
                if (p0 || false) {
                    var v8 int64 = 9223372036854775807
                    if ((1.5 <= 0.1) && (!(!((3.14159 / ((1e10 / float32((((0.1 * 1.5) + float64(float32(3.14159))) + 1.5))) - 1e10)) < -2.25)))) {
                        var v9 string = "a"
                        v10 := []int{fi5, (fi5 + fi5), fi5}
                        v11 := GenS1{F0: 252}
                        _ = v10
                        _ = v11
                        _ = v9
                    } else {
                        p0 = p0
                        var v12 uint8 = 161
                        _ = v12
                    }
                    var v13 int = (fi5 % ((127) | 1))
                    _ = v13
                    _ = v8
                } else {
                    var v14 int8 = -128
                    v15 := []int{(fi5 & fi5), (fi5 % ((865) | 1))}
                    _ = v14
                    _ = v15
                }
                var v16 int16 = -32768
                _ = v16
            } else {
                if p0 {
                    if (100 >= 42) {
                        v17 := []int{(fi5 & fi5)}
                        var v18 string = ""
                        v19 := GenS1{F0: 128}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    var v20 uint8 = 255
                    _ = v20
                } else {
                    var v21 string = ((((("\n\t" + "\n\t") + (("" + "foo bar") + ("foo bar" + ""))) + "foo bar") + (("" + "") + ("" + "a"))) + "")
                    v21 = v21
                    if (!(2147483648 == 2)) {
                        v22 := []int{}
                        if false {
                            p1 -= p1
                            fmt.Println(p1, fi5, p0)
                        } else {
                            switch (((v21 + v21) + v21) + "a") {
                            case "\n\t":
                                fmt.Println(p1, fi5, v21)
                            case "a":
                                p0 = true
                            }
                        }
                        _ = v22
                    }
                    _ = v21
                }
                p1 -= p1
            }
            _ = v7
        }
        var v23 string = ("a" + "\n\t")
        _ = fi5
        _ = v23
    }
    p1 += (p1 >> ((1) & 63))
    fmt.Println(p1)
    v24 := []int{36}
    v25 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    return 255
}

func main() {
    var v1 int16 = 1000
    v2 := []int{-2147483648, 851, -1, -9223372036854775808}
    v3 := make(map[string]int)
    var v26 float64 = 0.0
    v3["hello"] = 42
    fmt.Println(v1, v26)
    var v27 int64 = 9223372036854775807
    v26 = float64((1e10 / (((3.14159 * (0.1 * 0.0)) / -2.25) - (0.1 / (3.14159 * 1e10)))))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v26, v26)
    v27 += 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
