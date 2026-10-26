package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 uint16, p1 int8) int16 {
    switch ((("hello" + "") + "a") + "\n\t") {
    case "a":
        v6 := []int{32768}
        _ = v6
    case "":
        if false {
            var v7 string = ""
            _ = v7
        }
    case "hello":
        var v8 uint8 = 127
        _ = v8
    }
    var v9 float32 = 3.14159
    var v10 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 67
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 uint32 = 65535
    var v4 int = -363
    var v5 int16 = 1000
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            v3 = v3
            _ = fi12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint8 = 42
        fmt.Println(v14, v4, v4)
        _ = fi13
        _ = v14
    }
    v4 = 0
    var v15 uint64 = 1
    var v16 uint32 = v3
    var v17 uint8 = uint8(v15)
    fmt.Println(v15)
    var v18 int32 = 2147483647
    if true {
        if ("foo bar" < "hello") {
            var v19 bool = (true || true)
            switch (("hello" + ((("hello" + "") + "a") + "a")) + (("" + ("hello" + ("" + (("" + "\n\t") + "\n\t")))) + "\n\t")) {
            case "hello":
                var v20 int64 = -1
                _ = v20
            case "a":
                if ((3.14159 * (-2.25 / 0.1)) == 0.0) {
                    var v21 float32 = 0.0
                    var v22 uint32 = (v3 | v3)
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS1{F0: (^(v18)), F1: -9223372036854775808, F2: 785}
                    _ = v23
                }
            }
            if false {
                for fi24 := 0; fi24 < 15; fi24++ {
                    if false {
                        var v25 uint = 42
                        for fi26 := 0; fi26 < 13; fi26++ {
                            var v27 string = "hello"
                            _ = fi26
                            _ = v27
                        }
                        _ = v25
                    }
                    v28 := v1["a"]
                    for fi29 := 0; fi29 < 11; fi29++ {
                        var v30 uint8 = (v17 << ((0) & 63))
                        v31 := []int{(v4 >> ((492) & 63))}
                        if ((fi29 % ((9223372036854775807) | 1)) >= fi24) {
                            v32 := GenS2{F0: v18, F1: v18}
                            v19 = v19
                            _ = v32
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = fi24
                    _ = v28
                }
            }
            _ = v19
        }
        v33 := []int{}
        switch "\n\t" {
        case "foo bar":
            if false {
                _ = len(v2)
                var v34 int16 = v5
                _ = v34
            } else {
                if (!(true && true)) {
                    var v35 uint = 1
                    _ = v35
                }
            }
        case "hello":
            fmt.Println(v16)
        default:
            var v36 bool = ((v18 <= v18) && true)
            _ = v36
        }
        _ = v33
    } else {
        switch "" {
        case "hello":
            var v37 float32 = float32(3.14159)
            _ = v37
        case "\n\t":
            var v38 float64 = 1e10
            _ = v38
        case "":
            fmt.Println(v17)
        }
    }
    fmt.Println(v18, v4)
    v16 += (v16 + 65535)
    for fi39 := 0; fi39 < 6; fi39++ {
        if true {
            var v40 bool = false
            v2[-867] = int(v18)
            _ = v40
        }
        switch 9223372036854775807 {
        case 32768:
            v41 := v2[(v4 << ((529) & 63))]
            _ = v41
        case 951:
            v15 = (^(v15))
        case 4294967295:
            if false {
                if false {
                    if (4294967295 > 65535) {
                        v15 = 18446744073709551615
                    }
                    fmt.Println(fi39, v17)
                }
                var v42 int32 = (v18 % (((v18 | v18)) | 1))
                for fi43 := 0; fi43 < 13; fi43++ {
                    var v44 uint64 = (v15 >> ((42) & 63))
                    var v45 uint = 2
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                _ = v42
            }
        }
        _ = fi39
    }
    if ((^(v17)) >= uint8(v16)) {
        v16 += 4294967295
        var v46 uint16 = 0
        v17 -= (^(v17))
        _ = v46
    } else {
        v16 = (v16 % ((v16) | 1))
        fmt.Println(v17, v4, v18)
        v3 = v3
    }
    v5 = (v5 >> ((379) & 63))
    var v47 string = (("\n\t" + "a") + "foo bar")
    v4 += v4
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
