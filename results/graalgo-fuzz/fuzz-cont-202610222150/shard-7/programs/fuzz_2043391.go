package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int16, p1 uint64) uint16 {
    v6 := GenS1{F0: p1}
    switch ("hello" + "\n\t") {
    case "":
        if true {
            var v7 int = -75
            v8 := GenS2{F0: 42}
            v9 := []int{(v7 & v7), (v7 ^ (v7 * (-(v7))))}
            _ = v7
            _ = v8
            _ = v9
        }
    case "foo bar":
        switch -139 {
        case 127:
            for fi10 := 0; fi10 < 1; fi10++ {
                if false {
                    v11 := []int{(fi10 >> ((278) & 63))}
                    v11 = append(v11, (fi10 + fi10))
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if (!(false || (0.0 == (0.1 + float32((((1e10 * (0.0 - -2.25)) + ((1.5 / 1e10) / 1.5)) + (float64((0.1 + (((3.14159 + (3.14159 * 3.14159)) / -2.25) - -2.25))) - 1.5))))))) {
                            var v13 int64 = 9223372036854775807
                            var v14 uint = 4294967295
                            _ = v13
                            _ = v14
                        } else {
                            v15 := []int{42, (fi10 >> ((4294967295) & 63)), fi12, fi10}
                            p0 -= (p0 << ((255) & 63))
                            _ = v15
                        }
                        var v16 int16 = 77
                        _ = fi12
                        _ = v16
                    }
                }
                var v17 float32 = (0.1 * (0.0 + float32((1e10 + float64(0.0)))))
                _ = fi10
                _ = v17
            }
        }
    case "a":
        var v18 uint8 = 2
        _ = v18
    }
    if (!true) {
        switch -833 {
        case 127:
            var v19 bool = false
            _ = v19
        case 42:
            if (((!(!true)) && ((42 <= 1) && (!false))) && (0 == -1)) {
                for fi20 := 0; fi20 < 15; fi20++ {
                    fmt.Println(fi20)
                    var v21 float32 = 0.0
                    if true {
                        v21 -= v21
                        var v22 int8 = 100
                        _ = v22
                    }
                    _ = fi20
                    _ = v21
                }
                fmt.Println(p0, p1)
                p0 = (p0 & p0)
            } else {
                v23 := []int{-292, 2147483647, 32767, 704}
                if true {
                    switch 32768 {
                    case 127:
                        fmt.Println(p1)
                    default:
                        var v24 uint16 = 65535
                        _ = v24
                    }
                    var v25 uint32 = 4294967295
                    var v26 int64 = 1
                    _ = v25
                    _ = v26
                } else {
                    p1 = (^(p1))
                    fmt.Println(p1, p1)
                }
                _ = v23
            }
        default:
            var v27 uint64 = p1
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    v1 := []string{"a"}
    v2 := []int{1, 42, 32767, 127}
    var v3 bool = false
    var v4 string = "hello"
    var v5 uint16 = 835
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            fmt.Println(fi29, v4)
            v30 := GenS1{F0: 0}
            v3 = v3
            _ = fi29
            _ = v30
        }
        v31 := GenS2{F0: -9223372036854775808}
        if (false && ((false && v3) || v3)) {
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 int32 = 0
                fmt.Println(v5, v4)
                _ = fi32
                _ = v33
            }
        } else {
            v2 = append(v2, 873)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi28
        _ = v31
    }
    v4 = "foo bar"
    fmt.Println(v5, v5, v5)
    v3 = v3
    v4 = ((v4 + (("foo bar" + "a") + ((v4 + v4) + v4))) + "hello")
    switch v4 {
    case "a":
        if v3 {
            if len(v2) > 0 {
                v2[0] = 0
            }
            var v34 uint32 = 642
            _ = v34
        }
    case "hello":
        var v35 uint64 = 18446744073709551615
        _ = v35
    }
    v5 -= v5
    fmt.Println(v4, v5)
    var v36 float64 = (1.5 * 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
