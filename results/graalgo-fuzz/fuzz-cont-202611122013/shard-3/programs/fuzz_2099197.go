package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 204
    v3 := make(map[string]bool)
    var v4 float32 = -2.25
    var v5 uint32 = (^(v1))
    var v6 uint8 = (v2 ^ (v2 | v2))
    for fi7 := 0; fi7 < 7; fi7++ {
        if false {
            if true {
                v2 += 222
            } else {
                if (3.14159 <= 1.5) {
                    fmt.Println(v2, v6)
                    var v8 int16 = 1
                    _ = v8
                } else {
                    switch (("hello" + "a") + "hello") {
                    case "":
                        if ((v5 * v5) != v1) {
                            v9 := []int{(fi7 ^ fi7), (^(fi7)), -9223372036854775808, fi7}
                            v9 = append(v9, fi7)
                            _ = v9
                        }
                    default:
                        v10 := GenS1{F0: 42}
                        _ = v10
                    }
                    switch (-(fi7)) {
                    case 32767:
                        delete(v3, ((("" + "\n\t") + "foo bar") + "foo bar"))
                    }
                    switch -234 {
                    case -9223372036854775808:
                        v11 := GenS1{F0: 4294967295}
                        _ = v11
                    }
                }
            }
            var v12 int8 = 1
            _ = v12
        }
        var v13 uint8 = v2
        _ = fi7
        _ = v13
    }
    v14 := GenS1{F0: 18446744073709551615}
    v15 := GenS2{F0: false}
    if true {
        var v16 bool = true
        v4 = v4
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        switch ("\n\t" + "") {
        case "hello":
            v18 := GenS1{F0: 1}
            _ = v18
        case "a":
            v1 += (v1 & (^(v1)))
        case "foo bar":
            v19 := GenS1{F0: 18446744073709551615}
            _ = v19
        }
        switch ("\n\t" + ("foo bar" + "a")) {
        case "\n\t":
            v6 -= 10
        default:
            var v20 bool = true
            _ = v20
        }
        _ = fi17
    }
    v3["foo bar"] = ((false && true) && ((((v4 + v4) + 3.14159) / -2.25) >= (v4 / 3.14159)))
    var v21 float64 = 3.14159
    fmt.Println(v4)
    switch -1 {
    case 0:
        fmt.Println(v5)
    case 9223372036854775807:
        var v22 uint16 = 65535
        _ = v22
    case 932:
        v2 += (v6 >> ((621) & 63))
    default:
        v23 := GenS3{F0: 2}
        _ = v23
    }
    v24 := GenS3{F0: v6}
    for fi25 := 0; fi25 < 14; fi25++ {
        v4 -= v4
        var v26 int64 = -781
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 float64 = (v21 / (v21 - 0.1))
            v30 := GenS3{F0: v2}
            _ = fi28
            _ = v29
            _ = v30
        }
        v4 -= 0.1
        _ = fi27
    }
    if true {
        switch -69 {
        case 9223372036854775807:
            var v31 uint16 = 32767
            _ = v31
        case 0:
            var v32 uint64 = 0
            _ = v32
        }
        var v33 uint64 = 615
        switch (("\n\t" + "a") + "\n\t") {
        case "hello":
            v21 = 0.0
        case "a":
            v21 = v21
        case "foo bar":
            if false {
                v4 = float32(float64((v4 + ((float32(float64(float32(v21))) * float32(v21)) - v4))))
                if false {
                    v34 := []int{}
                    _ = v34
                }
                var v35 uint64 = v33
                _ = v35
            } else {
                v21 = (0.1 / (0.1 * -2.25))
                if (false && true) {
                    var v36 string = "\n\t"
                    _ = v36
                }
            }
        default:
            var v37 uint8 = 255
            _ = v37
        }
        _ = v33
    }
    var v38 uint8 = (v6 >> ((663) & 63))
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 bool = true
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
