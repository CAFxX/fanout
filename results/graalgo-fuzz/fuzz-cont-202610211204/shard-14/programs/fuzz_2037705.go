package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint, p1 uint8) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int8 = 127
            if (42 != 42) {
                v7 := []int{(fi5 - fi5), (fi5 * fi4)}
                _ = v7
            }
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    p1 -= (^(p1))
    switch "foo bar" {
    case "a":
        if false {
            var v8 uint32 = 2147483647
            _ = v8
        } else {
            var v9 string = "a"
            _ = v9
        }
    default:
        v10 := []int{}
        _ = v10
    }
    var v11 uint16 = 72
    _ = p0
    _ = p1
    _ = v11
    return (^(p0))
}

func main() {
    var v1 float64 = ((1e10 + float64((float32((1.5 - 3.14159)) + -2.25))) - (1e10 / float64(-2.25)))
    var v2 string = ((("\n\t" + "hello") + ("" + "")) + ("\n\t" + ""))
    var v3 uint = 756
    fmt.Println(v3, v1)
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int16 = 32767
        var v14 uint32 = 4294967295
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int64 = -1
    if ((v2 + (v2 + v2)) > v2) {
        v16 := GenS1{F0: 2}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                if (true && (!true)) {
                    var v19 bool = false
                    _ = v19
                } else {
                    v20 := GenS1{F0: 18}
                    if (false && (true || (v3 != (v3 / ((v3) | 1))))) {
                        v21 := []int{fi18}
                        _ = v21
                    }
                    _ = v20
                }
                var v22 int = fi17
                if true {
                    v3 = v3
                    switch v2 {
                    case "foo bar":
                        v23 := []int{957, fi17}
                        _ = v23
                    case "a":
                        var v24 uint = (v3 + v3)
                        _ = v24
                    }
                    v15 += (v15 * v15)
                    v25 := GenS2{F0: true}
                    if (fi18 == 4294967295) {
                        v15 -= -826
                        fmt.Println(v15)
                    }
                    _ = v25
                }
                _ = fi18
                _ = v22
            }
            v1 += v1
            _ = fi17
        }
        if false {
            var v26 bool = false
            switch -828 {
            case 2147483647:
                v3 -= genFn0(v3, 243)
            case 284:
                v27 := []int{-2147483648, 32767, 9223372036854775807, 2147483647}
                _ = v27
            }
            v28 := []int{61, 532, -470, 128}
            _ = v26
            _ = v28
        } else {
            var v29 float32 = 3.14159
            var v30 float32 = 0.1
            switch -351 {
            case 4294967295:
                var v31 int64 = 548
                _ = v31
            case -865:
                switch 42 {
                case 1:
                    switch v2 {
                    case "\n\t":
                        for fi32 := 0; fi32 < 17; fi32++ {
                            switch v2 {
                            case "\n\t":
                                v33 := GenS1{F0: 2}
                                _ = v33
                            }
                            _ = fi32
                        }
                    }
                default:
                    v2 = v2
                }
            }
            _ = v29
            _ = v30
        }
    }
    v34 := GenS2{F0: (!((42 != 0) && (0.0 < (0.1 + (-2.25 * (1e10 - -2.25))))))}
    fmt.Println(v15)
    fmt.Println(v2, v2)
    var v35 float64 = (0.0 / (float64((-2.25 - 0.1)) + (v1 - (0.1 * v1))))
    fmt.Println(v3, v1)
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 3; fi37++ {
            v1 = v35
            _ = fi37
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
