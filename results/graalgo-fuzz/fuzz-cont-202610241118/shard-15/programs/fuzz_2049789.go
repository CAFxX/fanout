package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32, p1 uint32, p2 int64) uint8 {
    if true {
        v7 := GenS1{F0: 127, F1: 18446744073709551615}
        _ = v7
    } else {
        switch -893 {
        case -505:
            p1 = (^(p1))
        }
    }
    v8 := []int{-586}
    var v9 int = 4294967295
    var v10 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 255
}
func genFn1() uint32 {
    var v11 uint32 = 110
    var v12 float32 = ((1e10 / 3.14159) - -2.25)
    switch ("" + ("a" + "a")) {
    case "a":
        if true {
            var v13 uint = 4294967295
            switch 4294967295 {
            case 444:
                v11 += (v11 | v11)
            case 128:
                v14 := []int{int(v13), 9223372036854775807}
                _ = v14
            }
            var v15 bool = (0 <= 100)
            _ = v13
            _ = v15
        } else {
            for fi16 := 0; fi16 < 16; fi16++ {
                for fi17 := 0; fi17 < 20; fi17++ {
                    var v18 uint = 18446744073709551615
                    v18 = (v18 * (v18 - (v18 << (((v18 * 42)) & 63))))
                    switch fi16 {
                    case 2147483647:
                        var v19 uint32 = uint32(v18)
                        _ = v19
                    }
                    _ = fi17
                    _ = v18
                }
                fmt.Println(v12, v11, v12)
                _ = fi16
            }
            if true {
                v20 := GenS1{F0: 127, F1: 442}
                fmt.Println(v11, v12)
                var v21 int8 = 43
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 1, F1: 0}
            _ = v22
        }
    }
    v12 = (3.14159 * -2.25)
    _ = v11
    _ = v12
    return (v11 ^ 0)
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int = 32767
    var v3 uint32 = 37
    v4 := make(map[string]int)
    var v5 int64 = 21
    var v6 string = "foo bar"
    if (!(v6 > ((((v6 + v6) + "foo bar") + ("" + v6)) + v6))) {
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS2{F0: 1000, F1: 96}
            v25 := []int{}
            for fi26 := 0; fi26 < 19; fi26++ {
                v3 = (v1 & 65535)
                delete(v4, v6)
                _ = fi26
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := []int{(fi27 - fi27), 32767, (^(fi27)), 595}
            var v29 string = v6
            for fi30 := 0; fi30 < 19; fi30++ {
                v3 = (v1 / ((65535) | 1))
                _ = fi30
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v31 := GenS2{F0: 42, F1: 21}
        _ = v31
    }
    switch v6 {
    case "a":
        fmt.Println(v1)
    case "foo bar":
        if false {
            var v32 int16 = 1
            _ = v32
        }
    default:
        if ((!false) && true) {
            v1 = (^(v3))
        } else {
            switch (^(v2)) {
            case 549:
                v33 := GenS1{F0: 1, F1: 712}
                _ = v33
            default:
                v6 = v6
            }
            if false {
                var v34 uint8 = 199
                fmt.Println(v5)
                _ = v34
            }
        }
    }
    var v35 uint16 = 336
    switch v6 {
    case "a":
        v1 = v1
    case "hello":
        switch v6 {
        case "foo bar":
            v6 = "\n\t"
        case "":
            delete(v4, "")
        case "\n\t":
            v36 := []int{v2, v2, 127, v2}
            _ = v36
        }
    case "\n\t":
        v6 = ((((v6 + v6) + (("foo bar" + "foo bar") + ("\n\t" + ((v6 + (v6 + ((((v6 + (v6 + "\n\t")) + v6) + "") + ""))) + "hello")))) + v6) + "\n\t")
    default:
        var v37 int64 = (v5 / ((42) | 1))
        _ = v37
    }
    fmt.Println(v3, v1, v5)
    var v38 string = v6
    v2 = v2
    var v39 float32 = (-2.25 - 0.1)
    var v40 uint16 = v35
    switch (v6 + v6) {
    case "\n\t":
        var v41 float64 = -2.25
        _ = v41
    case "":
        switch -491 {
        case -2147483648:
            v42 := GenS2{F0: 42, F1: 0}
            _ = v42
        case -1:
            v43 := []int{}
            _ = v43
        case 32768:
            if ((false || true) && true) {
                delete(v4, ("" + "a"))
            }
        default:
            fmt.Println(v1, v3)
        }
    case "hello":
        v4[v38] = v2
    }
    v44 := GenS1{F0: 42, F1: 1}
    var v45 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
