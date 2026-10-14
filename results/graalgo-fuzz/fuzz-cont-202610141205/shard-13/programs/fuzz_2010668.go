package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int, p1 int8, p2 int8) uint16 {
    v5 := GenS1{F0: 42}
    v6 := GenS1{F0: 2}
    for fi7 := 0; fi7 < 20; fi7++ {
        switch p0 {
        case 4294967295:
            v8 := GenS1{F0: 778}
            _ = v8
        case 127:
            v9 := GenS1{F0: 2}
            _ = v9
        }
        if false {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 int64 = 0
                switch -1 {
                case 9223372036854775807:
                    var v12 bool = false
                    _ = v12
                case 128:
                    v13 := GenS1{F0: 18446744073709551615}
                    v14 := []int{p0}
                    _ = v13
                    _ = v14
                }
                v11 += -9223372036854775808
                var v15 uint64 = 18446744073709551615
                if false {
                    p1 = 127
                }
                _ = fi10
                _ = v11
                _ = v15
            }
            fmt.Println(fi7, p0)
        }
        p0 -= (-(fi7))
        _ = fi7
    }
    switch p0 {
    case -2147483648:
        for fi16 := 0; fi16 < 8; fi16++ {
            p0 += 4294967295
            _ = fi16
        }
    case 0:
        var v17 string = "\n\t"
        _ = v17
    case -9223372036854775808:
        var v18 uint = 0
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]bool)
    var v3 int32 = -595
    var v4 uint8 = 128
    if (("hello" + "") != "foo bar") {
        var v19 int32 = (v3 | -360)
        _ = v19
    }
    v20 := GenS1{F0: 4294967295}
    if (((!((2 > 676) && true)) || (122 != v4)) || (543 > 1000)) {
        if true {
            var v21 uint64 = 1
            v22 := []int{-9223372036854775808}
            _ = len(v2)
            _ = v21
            _ = v22
        } else {
            v23 := GenS1{F0: 745}
            var v24 uint = v1
            switch "foo bar" {
            case "\n\t":
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := []int{fi25}
                    _ = fi25
                    _ = v26
                }
            }
            _ = v23
            _ = v24
        }
    }
    v3 = (v3 % (((v3 >> ((v1) & 63))) | 1))
    v27 := []int{2147483647, -9223372036854775808, -83, -552}
    switch -378 {
    case 4294967295:
        fmt.Println(v1)
    case 596:
        v1 += 163
    case 127:
        v28 := GenS1{F0: 18446744073709551615}
        _ = v28
    }
    switch -2147483648 {
    case -9223372036854775808:
        _ = len(v27)
    }
    v29 := v2[""]
    var v30 uint16 = 42
    if v29 {
        var v31 uint64 = uint64(v1)
        if false {
            var v32 int32 = v3
            var v33 int64 = -1
            var v34 uint8 = v4
            _ = v32
            _ = v33
            _ = v34
        } else {
            var v35 uint16 = v30
            switch 2147483647 {
            case -389:
                for fi36 := 0; fi36 < 3; fi36++ {
                    v29 = ((0.1 + (-2.25 - 1e10)) == (3.14159 + 0.0))
                    var v37 string = (((((("foo bar" + "\n\t") + "a") + ("\n\t" + "a")) + "a") + "foo bar") + "")
                    _ = fi36
                    _ = v37
                }
            case -1:
                if (!(1.5 != float64(float32((1.5 * ((3.14159 - 0.1) + ((((0.1 / 1.5) - float64((1.5 - 1e10))) * 0.1) * (1e10 - 1e10)))))))) {
                    v38 := GenS1{F0: (^(v31))}
                    _ = v38
                }
            case 2147483647:
                v31 -= 44
            default:
                if v29 {
                    var v39 uint32 = 2147483648
                    var v40 float64 = 1e10
                    _ = v39
                    _ = v40
                } else {
                    fmt.Println(v3, v1)
                    var v41 uint16 = 255
                    _ = v41
                }
            }
            _ = v35
        }
        v2[""] = (1 >= (v1 & uint(v30)))
        _ = v31
    }
    fmt.Println(v3, v4)
    delete(v2, "foo bar")
    v30 += 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
