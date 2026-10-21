package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0() int {
    if (!(1.5 < 0.1)) {
        var v6 float64 = 3.14159
        v6 -= v6
        var v7 int = 42
        _ = v6
        _ = v7
    }
    var v8 int16 = -118
    var v9 float32 = 1.5
    var v10 int64 = -1
    var v11 uint16 = 32767
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 540
}
func genFn1() uint {
    v12 := []int{}
    var v13 uint = 126
    _ = v12
    _ = v13
    return v13
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = (("hello" + (("a" + (("" + "hello") + ("foo bar" + "hello"))) + (("a" + "") + ((("a" + ("hello" + "")) + "") + "hello")))) + (("" + "") + "a"))
    var v3 int64 = 809
    var v4 int8 = 100
    v5 := make(map[string]bool)
    var v14 string = (v2 + ("foo bar" + (((v2 + ("a" + v2)) + ((v2 + ((v2 + "foo bar") + v2)) + (("foo bar" + ("\n\t" + "\n\t")) + ("" + "\n\t")))) + "")))
    if true {
        fmt.Println(v14, v14, v14)
    } else {
        delete(v5, v2)
        if (true && false) {
            var v15 int = -20
            v16 := []int{}
            _ = v15
            _ = v16
        } else {
            v17 := v5[v2]
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 float64 = -2.25
                var v20 uint16 = 910
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 1; fi21++ {
                switch v14 {
                case "\n\t":
                    v17 = false
                case "a":
                    v14 = v14
                }
                var v22 bool = false
                switch v14 {
                case "":
                    v17 = true
                case "foo bar":
                    delete(v1, 2)
                }
                _ = fi21
                _ = v22
            }
            _ = v17
        }
        v2 = ("a" + v14)
    }
    v3 = (-(v3))
    var v23 int = int(v4)
    if false {
        if (((v4 + 29) >= v4) || false) {
            v24 := []int{(-(v23)), (v23 ^ -2147483648), v23}
            _ = v24
        } else {
            var v25 int64 = (v3 / ((v3) | 1))
            _ = v25
        }
        var v26 uint64 = 0
        v27 := GenS1{F0: 65535, F1: (v4 ^ (v4 + int8(v3)))}
        _ = v26
        _ = v27
    }
    var v28 uint16 = 281
    switch v23 {
    case 2147483647:
        v29 := GenS2{F0: 65535}
        _ = v29
    case -2147483648:
        for fi30 := 0; fi30 < 13; fi30++ {
            for fi31 := 0; fi31 < 2; fi31++ {
                if false {
                    v32 := []int{genFn0(), 710}
                    _ = v32
                }
                v33 := GenS1{F0: 414, F1: v4}
                var v34 uint = 42
                _ = fi31
                _ = v33
                _ = v34
            }
            v35 := GenS1{F0: 2147483648, F1: 127}
            _ = fi30
            _ = v35
        }
    case 127:
        v4 = (v4 % ((v4) | 1))
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        fmt.Println(v28, v3, v14)
        v37 := GenS2{F0: v28}
        v38 := v1[359]
        _ = fi36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 int64 = 0
        _ = fi39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
