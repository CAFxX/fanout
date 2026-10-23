package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0(p0 bool, p1 int32) int8 {
    v4 := []int{}
    v5 := GenS1{F0: 1, F1: 42}
    if (!true) {
        v6 := GenS1{F0: 2, F1: 774}
        _ = v6
    } else {
        p1 = 0
        fmt.Println(p1)
        for fi7 := 0; fi7 < 8; fi7++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 bool = p0
                v9 = v9
                _ = fi8
                _ = v9
            }
            var v10 uint = 255
            _ = fi7
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 87
}

func main() {
    var v1 string = ("\n\t" + ("a" + ("foo bar" + ("a" + "hello"))))
    var v2 uint32 = 65535
    v3 := []int{-2147483648, 1, 32768}
    for fi11 := 0; fi11 < 10; fi11++ {
        _ = len(v3)
        fmt.Println(v1, v1, v2)
        var v12 uint64 = 0
        _ = fi11
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v13 := GenS1{F0: 255, F1: (v2 % ((2147483647) | 1))}
    fmt.Println(v2, v2)
    var v14 bool = false
    switch 128 {
    case 144:
        v15 := GenS1{F0: 42, F1: 4294967295}
        _ = v15
    case 128:
        var v16 bool = (!(42 >= -1))
        _ = v16
    }
    if false {
        if false {
            fmt.Println(v14)
        } else {
            _ = len(v3)
            v17 := GenS2{F0: v1, F1: 989}
            if (!((!(v1 <= (("\n\t" + v1) + "a"))) && ((v1 + v1) >= ((v1 + "foo bar") + "\n\t")))) {
                if v14 {
                    v2 = (v2 % (((v2 % (((v2 & (^(v2)))) | 1))) | 1))
                } else {
                    v14 = (v14 && (!(2 == 42)))
                }
                v1 = ((((("foo bar" + ("foo bar" + (("hello" + (((v1 + ((v1 + "foo bar") + (v1 + v1))) + (v1 + v1)) + "")) + v1))) + v1) + v1) + (((v1 + "hello") + v1) + "foo bar")) + ("hello" + ""))
            } else {
                if v14 {
                    if (v14 && (((42 < 1) && v14) || v14)) {
                        for fi18 := 0; fi18 < 17; fi18++ {
                            v19 := GenS3{F0: (fi18 > fi18), F1: -459}
                            var v20 string = ((v1 + v1) + (v1 + v1))
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v14 = false
                            _ = fi21
                        }
                    }
                }
                var v22 string = "a"
                v3 = append(v3, -46)
                _ = v22
            }
            _ = v17
        }
        fmt.Println(v14, v14, v1)
        var v23 int64 = -1
        _ = v23
    } else {
        v2 += 42
    }
    v2 -= 65535
    switch "hello" {
    case "foo bar":
        var v24 int = -295
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
