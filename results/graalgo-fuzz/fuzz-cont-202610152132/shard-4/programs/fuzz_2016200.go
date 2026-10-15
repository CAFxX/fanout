package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 uint8, p1 string, p2 uint32) string {
    fmt.Println(p0, p0)
    v6 := GenS1{F0: -128}
    v7 := []int{127, -615, 522, -450}
    for fi8 := 0; fi8 < 5; fi8++ {
        if true {
            switch 2147483647 {
            case 1:
                p0 = p0
            case 32767:
                v7 = append(v7, fi8)
            case 2147483647:
                if ("hello" > ((p1 + (p1 + p1)) + p1)) {
                    fmt.Println(p2, p2)
                    if (2 != 4294967295) {
                        var v9 uint64 = 23
                        _ = v9
                    }
                } else {
                    v10 := GenS1{F0: -1}
                    v11 := GenS2{F0: 2, F1: 285, F2: fi8}
                    _ = v10
                    _ = v11
                }
            default:
                var v12 bool = (!(-2.25 > (-2.25 * (1.5 + 1e10))))
                _ = v12
            }
            if ((!(!(-887 == -426))) && false) {
                v13 := GenS1{F0: 42}
                var v14 uint16 = 1
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v13
                _ = v14
            } else {
                v15 := GenS1{F0: 44}
                _ = v15
            }
        } else {
            if true {
                if true {
                    var v16 uint32 = p2
                    var v17 uint32 = (v16 & (p2 - (^(p2))))
                    var v18 uint16 = 42
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS2{F0: 0, F1: 1000, F2: 42}
                    _ = v19
                }
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 string, p1 string) int8 {
    if false {
        if false {
            fmt.Println(p0, p0)
            for fi20 := 0; fi20 < 15; fi20++ {
                if true {
                    v21 := []int{(-(fi20))}
                    _ = len(v21)
                    _ = v21
                }
                p1 = ""
                _ = fi20
            }
        }
    } else {
        p1 = p1
        v22 := GenS3{F0: 206, F1: -1}
        p1 = p0
        _ = v22
    }
    p1 = "\n\t"
    v23 := GenS4{F0: ((p0 + (p1 + (("a" + p1) + (((p1 + p0) + "") + p1)))) + p0)}
    if false {
        if ((255 == 918) || true) {
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 uint64 = 4294967295
                if false {
                    v26 := []int{}
                    fmt.Println(v25)
                    _ = v26
                }
                var v27 int16 = -632
                _ = fi24
                _ = v25
                _ = v27
            }
        }
        if true {
            p0 = ((("" + p0) + (p1 + "foo bar")) + p1)
        }
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 int16 = -32768
            _ = fi28
            _ = v29
        }
    }
    _ = p0
    _ = p1
    _ = v23
    return 42
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    fmt.Println(v1)
    fmt.Println(v1)
    for fi30 := 0; fi30 < 7; fi30++ {
        v1 -= 18446744073709551615
        if false {
            if false {
                v1 += v1
                var v31 int = fi30
                switch "\n\t" {
                case "foo bar":
                    v32 := []int{42, fi30, fi30, (fi30 + -2147483648)}
                    _ = v32
                }
                _ = v31
            }
        }
        _ = fi30
    }
    switch -9223372036854775808 {
    case 1:
        switch 41 {
        case 258:
            v1 = v1
        }
    }
    var v33 int32 = 1
    v1 = (v1 << ((780) & 63))
    delete(v5, ((("a" + ("a" + "\n\t")) + "hello") + "foo bar"))
    v1 = v1
    var v34 int8 = 37
    fmt.Println(v33)
    var v35 float64 = (3.14159 - 3.14159)
    v33 += (v33 % ((-203) | 1))
    delete(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
