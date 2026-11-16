package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 int) bool {
    fmt.Println(p0)
    var v4 uint16 = 65535
    if (104 >= 0) {
        var v5 int = (p0 - 844)
        v6 := []int{}
        _ = v5
        _ = v6
    } else {
        v4 += (^(v4))
    }
    v7 := []int{p0, p0, p0, p0}
    _ = p0
    _ = v4
    _ = v7
    return false
}
func genFn1() uint16 {
    switch ("foo bar" + "hello") {
    case "a":
        var v8 float32 = ((3.14159 + float32(1.5)) + 1e10)
        _ = v8
    case "foo bar":
        var v9 string = ("\n\t" + (((("a" + ("hello" + ((("" + "a") + ("\n\t" + "")) + (("\n\t" + "foo bar") + ("" + ("\n\t" + ("\n\t" + "\n\t"))))))) + "") + "foo bar") + "foo bar"))
        _ = v9
    default:
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float32 = 0.0
            _ = fi10
            _ = v11
        }
    }
    var v12 int = -2147483648
    var v13 bool = ((42 == 100) && (-1 == 42))
    var v14 uint16 = 2
    v15 := []int{}
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 281
}

func main() {
    v1 := []string{"foo bar", "", "a", ""}
    var v2 int16 = 42
    v3 := make(map[string]string)
    v2 = 1000
    v16 := GenS1{F0: v2}
    fmt.Println(v2)
    if true {
        if ((1e10 * (3.14159 - float64((0.0 + (0.1 + 0.1))))) == (1.5 + 0.1)) {
            var v17 uint32 = 485
            v2 = (v2 | (v2 * v2))
            fmt.Println(v17)
            _ = v17
        }
    } else {
        v2 = v2
        var v18 int8 = 1
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int32 = -230
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v2 -= (v2 + (v2 % (((v2 ^ 472)) | 1)))
    for fi21 := 0; fi21 < 2; fi21++ {
        _ = len(v3)
        if false {
            var v22 int8 = -1
            v23 := GenS1{F0: (v2 % ((v2) | 1))}
            _ = v22
            _ = v23
        }
        fmt.Println(v2)
        _ = fi21
    }
    v2 = v2
    fmt.Println(v2)
    v24 := GenS1{F0: (v2 ^ v2)}
    switch 32768 {
    case -842:
        var v25 uint16 = 0
        _ = v25
    case -410:
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 bool = (!false)
            v3["a"] = "\n\t"
            v28 := GenS1{F0: v2}
            _ = fi26
            _ = v27
            _ = v28
        }
    case -1:
        v29 := GenS1{F0: (v2 + v2)}
        _ = v29
    default:
        switch 32768 {
        case 379:
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 uint = 692
                var v32 int64 = 1
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        v34 := GenS1{F0: (v2 - 1)}
        switch (fi33 + (fi33 << ((169) & 63))) {
        case 127:
            v2 += (v2 >> ((255) & 63))
        case 13:
            switch "hello" {
            case "foo bar":
                v35 := GenS2{F0: 2}
                _ = v35
            }
        case 228:
            for fi36 := 0; fi36 < 16; fi36++ {
                var v37 int = -1
                v38 := GenS2{F0: 2}
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        _ = fi33
        _ = v34
    }
    delete(v3, ("a" + ""))
    if true {
        var v39 uint64 = 4294967295
        _ = v39
    }
    switch 32767 {
    case -252:
        v2 += (v2 - -18)
    case 128:
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    case -9223372036854775808:
        var v40 int64 = -1
        _ = v40
    default:
        var v41 int = -295
        _ = v41
    }
    switch "hello" {
    case "a":
        var v42 uint8 = 181
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
}
