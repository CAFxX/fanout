package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8) int {
    switch "a" {
    case "a":
        p0 = p0
    case "foo bar":
        p0 = p0
    }
    p0 -= 226
    v5 := GenS1{F0: -577}
    p0 -= 42
    p0 = 166
    _ = p0
    _ = v5
    return 591
}
func genFn1(p0 string, p1 uint16) uint64 {
    v6 := []int{42}
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 = (("foo bar" + p0) + (p0 + p0))
        _ = fi7
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 196
    var v3 bool = true
    v4 := make(map[string]uint32)
    fmt.Println(v1, v2, v3)
    v8 := []int{5, 127, -1, 128}
    v2 = 1
    var v9 int = 128
    for fi10 := 0; fi10 < 12; fi10++ {
        if ((^(v1)) < (v1 * v1)) {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        v9 = (v9 - 1)
        _ = fi10
    }
    v11 := GenS1{F0: -260}
    fmt.Println(v3)
    _ = len(v8)
    fmt.Println(v1, v9, v9)
    for fi12 := 0; fi12 < 17; fi12++ {
        if false {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int = v9
                _ = fi13
                _ = v14
            }
        } else {
            v15 := GenS1{F0: 1}
            var v16 int32 = 1
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int = 603
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
        fmt.Println(v3, v3)
        v4["foo bar"] = 995
        _ = fi12
    }
    var v19 uint = 18446744073709551615
    v20 := GenS1{F0: -1}
    fmt.Println(v1)
    v21 := GenS2{F0: -911}
    var v22 uint32 = (v1 << ((v19) & 63))
    switch ((("\n\t" + "\n\t") + "foo bar") + (("hello" + "foo bar") + "hello")) {
    case "foo bar":
        v23 := GenS1{F0: -9223372036854775808}
        _ = v23
    case "a":
        if true {
            v2 = v2
            if v3 {
                v1 += (v1 ^ v1)
            }
        } else {
            var v24 uint64 = 897
            var v25 uint64 = uint64(v1)
            var v26 int16 = 0
            _ = v24
            _ = v25
            _ = v26
        }
    default:
        switch (v9 | (v9 + v9)) {
        case -977:
            switch (("foo bar" + "a") + "") {
            case "\n\t":
                var v27 float32 = ((0.0 / ((3.14159 + 0.1) - 3.14159)) / 3.14159)
                _ = v27
            case "hello":
                for fi28 := 0; fi28 < 18; fi28++ {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        delete(v4, "a")
                        _ = fi29
                    }
                    _ = fi28
                }
            case "foo bar":
                v30 := v4["foo bar"]
                _ = v30
            }
        case -779:
            v31 := GenS1{F0: 740}
            _ = v31
        default:
            var v32 bool = false
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        if true {
            switch ("\n\t" + "hello") {
            case "hello":
                var v34 uint16 = 323
                _ = v34
            case "a":
                if len(v8) > 0 {
                    v8[0] = (^(fi33))
                }
            default:
                for fi35 := 0; fi35 < 4; fi35++ {
                    var v36 uint = uint(v22)
                    switch (fi33 << (((v36 >> (((v36 & (v19 % (((v36 | 227)) | 1)))) & 63))) & 63)) {
                    case 0:
                        var v37 uint64 = 4294967295
                        _ = v37
                    default:
                        v4["a"] = 944
                    }
                    _ = fi35
                    _ = v36
                }
            }
        } else {
            v19 -= (v19 & 959)
        }
        v38 := GenS2{F0: 1}
        for fi39 := 0; fi39 < 20; fi39++ {
            fmt.Println(v2, v2, v9)
            v40 := GenS1{F0: 9223372036854775807}
            if (!(v3 && ((^(v1)) <= 65535))) {
                switch ("a" + "foo bar") {
                case "":
                    var v41 int32 = -58
                    _ = v41
                default:
                    fmt.Println(v9, v1, fi39)
                }
            }
            _ = fi39
            _ = v40
        }
        _ = fi33
        _ = v38
    }
    var v42 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v8
    _ = v9
}
